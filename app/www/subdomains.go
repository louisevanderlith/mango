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

var subdomains Subdomains

func init() {
	subdomains = make(Subdomains)
}

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
		rawURL, err := util.GetServiceURL(v.Name)

		if rawURL != "" && err == nil {
			vshost, err := url.Parse(rawURL)

			if err != nil {
				log.Print(err)
			}

			proxy := httputil.NewSingleHostReverseProxy(vshost)

			domainMux := http.NewServeMux()
			domainMux.HandleFunc("/", domainHandler(proxy))

			subdomains[v.Address] = domainMux
		} else {
			log.Printf("Skipping %s", v.Name)
			log.Print(err)
		}
	}
}

func defaultMuxSetup() {
	fs := http.FileServer(http.Dir("static"))

	defaultMux := http.NewServeMux()
	defaultMux.Handle("/static/", http.StripPrefix("/static/", fs))
	defaultMux.Handle("/", fs)

	subdomains["www"] = defaultMux
}

func domainHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}
