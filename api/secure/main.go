package main

import (
	_ "github.com/louisevanderlith/mango/api/secure/routers"

	"github.com/astaxie/beego"
)

var instanceKey string

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	
	// Register with router
	srv := logic.Service{
		Environment: beego.BConfig.RunMode,
		Name:        beego.BConfig.AppName,
		URL:         "http://localhost:" + beego.AppConfig.String("httpport"),
		Type:        "service"
	}

	key, err := logic.Register()

	if err != nil{
		log.Panic(err)
	} else {
		instanceKey = key
		logic.BuildDatabase(instanceKey)
		beego.Run()
	}
}
