package main

import (
	"log"

	"github.com/louisevanderlith/mango/api/artifact/routers"
	_ "github.com/louisevanderlith/mango/core/artifact"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"

	"github.com/astaxie/beego"
)

func main() {
	mode := beego.BConfig.RunMode

	if mode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Register with router
	name := beego.BConfig.AppName
	srv := util.NewService(mode, enums.API)

	port := beego.AppConfig.String("httpport")
	instKey, err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(instKey)
		beego.Run()
	}
}
