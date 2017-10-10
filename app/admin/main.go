package main

import (
	_ "github.com/louisevanderlith/mango/app/admin/routers"
	"github.com/astaxie/beego"
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
	key, err := util.Register(srv, discURL, port)

	if err != nil {
		log.Print(err)
	} else {
		instanceKey = key
		beego.Run()
	}
}

