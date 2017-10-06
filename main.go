package main

import (
	"log"
	"net/http"

	"github.com/astaxie/beego"

	"github.com/louisevanderlith/mango/db/secure"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

var (
	instanceKey string
	subdomains  Subdomains
)

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
	registerSubdomains(discURL)

	log.Println("Listening...")
	err := http.ListenAndServe(beego.AppConfig.String("httpport"), subdomains)

	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}
}
