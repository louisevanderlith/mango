package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/louisevanderlith/mango/util"
)

type Subdomains map[string]http.Handler

func (subdomains Subdomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	domainParts := strings.Split(r.Host, ".")
	mux := subdomains[domainParts[0]]

	if mux == nil {
		// if the subdomain is not found, render default www website
		mux = subdomains["www"]
	}

	mux.ServeHTTP(w, r)
}

func registerSubdomains(discURL string) {
	defaultMuxSetup()

	domains := loadSettings()

	for _, v := range *domains {
		rawURL, err := util.GetServiceURL(instanceKey, v.Name, discURL)
		vshost, err := url.Parse(rawURL)

		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(vshost)

		domainMux := http.NewServeMux()
		domainMux.HandleFunc("/", domainHandler(proxy))

		subdomains[v.Address] = domainMux
	}
}

func defaultMuxSetup() {
	fs := http.FileServer(http.Dir("web"))

	defaultMux := http.NewServeMux()
	defaultMux.Handle("/web/", http.StripPrefix("/web/", fs))
	defaultMux.Handle("/", fs)

	subdomains["www"] = defaultMux
}

func domainHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("PROXY Activated...")
		p.ServeHTTP(w, r)
	}
}
