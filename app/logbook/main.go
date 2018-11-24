package main

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/logbook/routers"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

func main() {
	mode := beego.BConfig.RunMode
	appName := beego.BConfig.AppName

	// Register with router
	srv := util.NewService(mode, appName, enums.APP)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
