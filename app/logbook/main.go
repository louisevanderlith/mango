package main

import (
	"log"

	"github.com/astaxie/beego"
	_ "github.com/louisevanderlith/mango/app/logbook/routers"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

var instanceKey string

func main() {
	// Register with router
	srv := util.Service{
		Environment: enums.GetEnvironment(beego.BConfig.RunMode),
		Name:        beego.BConfig.AppName,
		Type:        enums.APP}

	discURL := beego.AppConfig.String("discovery")
	port := beego.AppConfig.String("httpport")
	key, err := srv.Register(discURL, port)

	if err != nil {
		log.Print(err)
	} else {
		instanceKey = key
		beego.Run()
	}
}
