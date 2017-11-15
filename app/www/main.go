package main

import (
	"log"
	"net/http"

	"github.com/astaxie/beego"

	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

func main() {
	// Register with router
	srv := util.Service{
		Environment: enums.GetEnvironment(beego.AppConfig.String("runmode")),
		Name:        beego.AppConfig.String("appname"),
		Type:        enums.APP}

	port := beego.AppConfig.String("httpport")
	_, err := srv.Register(port)

	if err != nil {
		log.Print(err)
	} else {
		setupHost(port)
	}
}

func setupHost(port string) {
	registerSubdomains()

	log.Println("Listening...")
	err := http.ListenAndServe(":"+port, subdomains)

	if err != nil {
		log.Print("ListenAndServe: ", err)
	}
}
