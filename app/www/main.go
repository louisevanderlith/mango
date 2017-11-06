package main

import (
	"log"
	"net/http"

	"github.com/astaxie/beego"

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
	port := beego.AppConfig.String("httpport")
	key, err := srv.Register(discURL, port)

	if err != nil {
		log.Print(err)
	} else {
		instanceKey = key
		setupHost(discURL, port)
	}
}

func setupHost(discURL, port string) {
	registerSubdomains(discURL)

	log.Println("Listening...")
	err := http.ListenAndServe(":"+port, subdomains)

	if err != nil {
		log.Print("ListenAndServe: ", err)
	}
}
