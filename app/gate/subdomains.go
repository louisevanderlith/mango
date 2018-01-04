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

	handleSession(r.URL, w)

	// CertBot requires tests on well-known for SSL Certs
	if !strings.Contains(r.URL.String(), "well-known") {
		domainParts := strings.Split(r.Host, ".")
		result = getMux(subdomains, domainParts)
	}

	result.ServeHTTP(w, r)
}

func handleSession(url *url.URL, w http.ResponseWriter) {
	if strings.Contains(url.String(), "?token=") {
		sessionID := url.Query().Get("token")
		url.Path = removeToken(url.String())

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

func removeToken(url string) string {
	var result string
	idx := strings.LastIndex(url, "?token")

	if idx != -1 {
		result = url[:idx]
	}

	return result
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
		rawURL, err := util.GetServiceURL(v.Name, false)

		if rawURL != "" && err == nil {
			vshost, err := url.Parse(rawURL)

			if err != nil {
				log.Printf("registerSubdomains: %s", err)
			}

			proxy := httputil.NewSingleHostReverseProxy(vshost)

			domainMux := http.NewServeMux()
			domainMux.HandleFunc("/", domainHandler(proxy))

			subdomains[v.Address] = domainMux
			log.Print(v.Name, " ", v.Address, " ", rawURL)
		} else {
			log.Printf("registerSubdomains: %s", err)
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
