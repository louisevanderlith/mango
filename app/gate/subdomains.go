package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/util"
)

type Subdomains map[string]http.Handler

var subdomains Subdomains

func init() {
	subdomains = make(Subdomains)
}

func (subdomains Subdomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result := subdomains["ssl"]

	// CertBot requires tests on well-known for SSL Certs
	if !strings.Contains(r.URL.String(), "well-known") {
		handleSession(*r.URL, w)
		domainParts := strings.Split(r.Host, ".")
		result = getMux(subdomains, domainParts)
	}

	result.ServeHTTP(w, r)
}

func handleSession(url url.URL, w http.ResponseWriter) {
	if strings.Contains(url.String(), "?token=") {
		sessionID := url.Query().Get("token")

		if sessionID != "" {
			cookie := http.Cookie{
				Name:     "avosession",
				Path:     "/",
				Value:    sessionID,
				HttpOnly: true,
				MaxAge:   0,
			}

			http.SetCookie(w, &cookie)
		}
	}
}

func getMux(subdomains Subdomains, domainParts []string) http.Handler {
	result := subdomains["www"]

	webMux, webOK := websiteMux(subdomains, domainParts)
	if webOK {
		result = webMux
	} else {
		subMux, subOk := subdomainMux(subdomains, domainParts)
		if subOk {
			result = subMux
		}
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
			subdomainMuxSetup(v)
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

	subdomains["ssl"] = sslMux
}

func staticMuxSetup(setting DomainSetting) {
	statMux := http.NewServeMux()
	fullPath := "/static/" + setting.Name + "/"
	fs := http.FileServer(http.FileSystem(http.Dir(fullPath)))
	statMux.Handle(fullPath, http.StripPrefix(fullPath, fs))
	statMux.Handle("/", fs)

	subdomains[setting.Address] = statMux
}

func subdomainMuxSetup(setting DomainSetting) {
	rawURL, err := util.GetServiceURL(setting.Name, false)

	if rawURL != "" && err == nil {
		vshost, err := url.Parse(rawURL)

		if err != nil {
			log.Printf("subdomainMuxSetup: %s", err)
		}

		proxy := httputil.NewSingleHostReverseProxy(vshost)

		domainMux := http.NewServeMux()
		domainMux.HandleFunc("/", domainHandler(proxy))

		subdomains[setting.Address] = domainMux
		log.Print(setting.Name, " ", setting.Address, " ", rawURL)
	} else {
		log.Printf("subdomainMuxSetup: %s", err)
	}
}

func domainHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}
