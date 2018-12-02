package main

import (
	"github.com/louisevanderlith/mango/api/router/routers"
	"github.com/louisevanderlith/mango/pkg"
	"github.com/louisevanderlith/mango/pkg/enums"

	"github.com/astaxie/beego"
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

	routers.Setup(srv)
	beego.Run()
}
