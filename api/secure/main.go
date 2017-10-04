package main

import (
	"log"

	_ "github.com/louisevanderlith/mango/api/secure/routers"
	"github.com/louisevanderlith/mango/db/secure"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"

	"github.com/astaxie/beego"
)

var instanceKey string

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Register with router
	srv := util.Service{
		Environment: enums.GetEnvironment(beego.BConfig.RunMode),
		Name:        beego.BConfig.AppName,
		Type:        enums.API}

	discURL := beego.AppConfig.String("discovery")
	key, err := util.Register(srv, discURL)

	if err != nil {
		log.Panic(err)
	} else {
		instanceKey = key
		secure.NewDatabase(instanceKey, discURL)
		beego.Run()
	}
}
