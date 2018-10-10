package domains

import (
	"log"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
)

type Subdomains struct {
	subs map[string]http.Handler
}

func NewSubdomains() *Subdomains {
	subs := make(map[string]http.Handler)
	return &Subdomains{subs}
}

const (
	token = "token"
	ssl   = "ssl"
	www   = "www"
)

func (s *Subdomains) Add(name string, handler http.Handler) {
	s.subs[name] = handler
}

func RegisterSubdomains(instanceID string) *Subdomains {
	result := NewSubdomains()
	result.Add(ssl, sslMuxSetup())

	confDomains := loadSettings()

	for _, v := range *confDomains {
		handl, err := v.SetupMux(instanceID)

		if err != nil {
			log.Printf("Register Subdomains: %s - %s\n", v.Name, err.Error())
		}

		result.Add(v.Address, handl)
	}

	return result
}

func (d *Subdomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// CertBot requires tests on well-known for SSL Certs
	if strings.Contains(r.URL.String(), "well-known") {
		sslHand, _ := d.subs[ssl]

		sslHand.ServeHTTP(w, r)
		return
	}

	handleSession(*r.URL, w)
	domainParts := strings.Split(r.Host, ".")
	sdomainName := domainParts[0]

	result := d.GetMux(sdomainName)

	result.ServeHTTP(w, r)
}

func (d *Subdomains) GetMux(subdomain string) http.Handler {
	result, ok := d.subs[subdomain]

	if !ok {
		return d.subs[www]
	}

	return result
}

func sslMuxSetup() http.Handler {
	sslMux := http.NewServeMux()
	certPath := beego.AppConfig.String("certpath")
	fullCertPath := http.FileSystem(http.Dir(certPath))
	fs := http.FileServer(fullCertPath)
	challengePath := "/.well-known/acme-challenge/"

	sslMux.Handle(challengePath, fs)

	return sslMux
}
