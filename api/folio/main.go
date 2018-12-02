package main

import (
	"log"

	"github.com/louisevanderlith/mango/api/folio/routers"
	"github.com/louisevanderlith/mango/pkg"
	"github.com/louisevanderlith/mango/pkg/enums"

	"github.com/astaxie/beego"
	_ "github.com/louisevanderlith/mango/core/folio"
)

func main() {
	mode := beego.BConfig.RunMode

	if mode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Register with router
	appName := beego.BConfig.AppName
	srv := util.NewService(mode, appName, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
