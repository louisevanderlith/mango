package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/astaxie/beego"

	"github.com/louisevanderlith/mango/db/secure"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

var instanceKey string

func main() {
	// Register with router
	srv := util.Service{
		Environment: enums.GetEnvironment(beego.AppConfig.String("runmode")),
		Name:        beego.AppConfig.String("appname"),
		Type:        enums.PROXY}

	discURL := beego.AppConfig.String("discovery")
	key, err := util.Register(srv, discURL)

	if err != nil {
		log.Panic(err)
	} else {
		instanceKey = key
		secure.NewDatabase(instanceKey, discURL)
		setupHost(discURL)
	}
}

func setupHost(discURL string) {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))
	http.Handle("/", fs)

	registerSubdomains(discURL)

	log.Println("Listening...")
	err := http.ListenAndServe(beego.AppConfig.String("httpport"), nil)

	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}
}

func registerSubdomains(discURL string) {
	domains := loadSettings()

	for _, v := range *domains {
		rawURL, err := util.GetServiceURL(instanceKey, v.Name, discURL)

		vshost, err := url.Parse(rawURL)
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(vshost)
		http.HandleFunc(v.Address, domainHandler(proxy))
	}
}

func domainHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	log.Print("PROXY doing it's stuff. More logging to come")

	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}
