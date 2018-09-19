package domains

import (
	"net/http"
	"net/http/httputil"
	"strings"
)

type Subdomains map[string]http.Handler

const (
	token = "token"
	ssl   = "ssl"
	www   = "www"
)

var domains Subdomains

func init() {
	domains = make(Subdomains)
}

func (d Subdomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result := d[ssl]

	// CertBot requires tests on well-known for SSL Certs
	if !strings.Contains(r.URL.String(), "well-known") {
		handleSession(*r.URL, w)
		domainParts := strings.Split(r.Host, ".")
		result = getMux(d, domainParts)
	}

	result.ServeHTTP(w, r)
}

func domainHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}
