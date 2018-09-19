package domains

import (
	"log"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
)

func getMux(subdomains Subdomains, domainParts []string) http.Handler {
	result := subdomains[www]

	webMux, webOK := websiteMux(subdomains, domainParts)
	if webOK {
		return webMux
	}

	subMux, subOk := subdomainMux(subdomains, domainParts)
	if subOk {
		return subMux
	}

	return result
}

func subdomainMux(subdomains Subdomains, domainParts []string) (http.Handler, bool) {
	result, ok := subdomains[domainParts[0]]

	return result, ok
}

func websiteMux(subdomains Subdomains, domainParts []string) (http.Handler, bool) {
	hostParts := domainParts[1:]
	host := strings.Join(hostParts, ".")

	result, ok := subdomains[host]

	return result, ok
}

func registerSubdomains() {
	sslMuxSetup()

	domains := loadSettings()

	for _, v := range *domains {
		if v.Type == "Subdomain" {
			err := subdomainMuxSetup(v)

			if err != nil {
				log.Println(err)
			}
		}

		if v.Type == "Static" {
			staticMuxSetup(v)
		}
	}
}

func sslMuxSetup() {
	sslMux := http.NewServeMux()
	certPath := beego.AppConfig.String("certpath")
	fullCertPath := http.FileSystem(http.Dir(certPath))
	fs := http.FileServer(fullCertPath)
	challengePath := "/.well-known/acme-challenge/"

	sslMux.Handle(challengePath, fs)

	domains[ssl] = sslMux
}
