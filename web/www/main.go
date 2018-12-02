package main

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/www/routers"
	"github.com/louisevanderlith/mango/pkg"
	"github.com/louisevanderlith/mango/pkg/enums"
)

func main() {
	// Register with router
	mode := beego.BConfig.RunMode
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
