package main

import (
	"log"

	_ "github.com/louisevanderlith/mango/core/auto"
	"github.com/louisevanderlith/mango/util/enums"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/auto/routers"
	"github.com/louisevanderlith/mango/util"
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
