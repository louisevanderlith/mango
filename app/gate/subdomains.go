package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/louisevanderlith/mango/util"
	"github.com/astaxie/beego"
)

type Subdomains map[string]http.Handler

var subdomains Subdomains

func init() {
	subdomains = make(Subdomains)
}

func (subdomains Subdomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result := subdomains["www"]

	// CertBot requires tests on well-known for SSL Certs
	if !strings.Contains(r.URL.String(), "well-known") {
		domainParts := strings.Split(r.Host, ".")

		result = getMux(subdomains, domainParts)
	}

	result.ServeHTTP(w, r)
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
	hostParts := remove(domainParts, 0)
	host := strings.Join(hostParts, ".")

	result, ok := subdomains[host]

	return result, ok
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func registerSubdomains() {
	sslMuxSetup()

	domains := loadSettings()

	for _, v := range *domains {
		rawURL, err := util.GetServiceURL(v.Name)

		if rawURL != "" && err == nil {
			vshost, err := url.Parse(rawURL)

			if err != nil {
				log.Printf("registerSubdomains: ", err)
			}

			proxy := httputil.NewSingleHostReverseProxy(vshost)

			domainMux := http.NewServeMux()
			domainMux.HandleFunc("/", domainHandler(proxy))

			subdomains[v.Address] = domainMux
		} else {
			log.Printf("Skipping %s", v.Name)
			log.Printf("registerSubdomains: ", err)
		}
	}
}

func sslMuxSetup() {
	sslMux := http.NewServeMux()
	certPath := beego.AppConfig.String("certpath")
	fs := http.FileServer(http.FileSystem(http.Dir(certPath)))
	sslMux.Handle("/.well-known/acme-challenge/", fs)

	subdomains["ssl"] = sslMux
}

func domainHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}
