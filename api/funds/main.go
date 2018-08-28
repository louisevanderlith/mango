package main

import (
	"log"

	"github.com/astaxie/beego"
	_ "github.com/louisevanderlith/mango/core/funds"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

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

	port := beego.AppConfig.String("httpport")
	_, err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		beego.Run()
	}
}
