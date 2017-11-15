package main

import (
	"log"
	"net/http"

	"github.com/astaxie/beego"

	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
	"golang.org/x/net/http2"
	"fmt"
)

func main() {
	// Register with router
	srv := util.Service{
		Environment: enums.GetEnvironment(beego.AppConfig.String("runmode")),
		Name:        beego.AppConfig.String("appname"),
		Type:        enums.APP}

	httpsPort := beego.AppConfig.String("httpsport")

	_, err := srv.Register(httpsPort)

	if err != nil {
		log.Printf("Register: ", err)
	} else {
		httpPort := beego.AppConfig.String("httpport")
		setupHost(httpPort, httpsPort)
	}
}

func setupHost(httpPort, httpsPort string) {
	registerSubdomains()

	var srv http.Server
	srv.Addr = ":" + httpsPort
	srv.Handler = subdomains

	cerr := http2.ConfigureServer(&srv, nil)

	if cerr != nil {
		log.Printf("ConfigureServer: ", cerr)
	}

	log.Println("Listening...")

	go srv.ListenAndServeTLS("host.cert", "host.key")

	err := http.ListenAndServe(":" + httpPort, http.HandlerFunc(redirectTLS))

	if err != nil {
		log.Print("ListenAndServe: ", err)
	}
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	moveURL := fmt.Sprintf("https://%s%s", r.Host, r.RequestURI)
	http.Redirect(w, r, moveURL, http.StatusTemporaryRedirect)
}
