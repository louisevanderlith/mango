package main

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/admin/routers"
	"github.com/louisevanderlith/mango/pkg"
	"github.com/louisevanderlith/mango/pkg/enums"
)

func main() {
	mode := beego.BConfig.RunMode

	// Register with router
	name := beego.BConfig.AppName
	srv := util.NewService(mode, name, enums.APP)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
