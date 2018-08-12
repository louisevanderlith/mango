package main

import (
	"log"

	_ "github.com/louisevanderlith/mango/core/auto"
	"github.com/louisevanderlith/mango/util/enums"

	"github.com/astaxie/beego"
	_ "github.com/louisevanderlith/mango/app/auto/routers"
	"github.com/louisevanderlith/mango/util"
)

func main() {
	// Register with router
	srv := util.Service{
		Environment: enums.GetEnvironment(beego.AppConfig.String("runmode")),
		Name:        beego.AppConfig.String("appname"),
		Type:        enums.APP}

	port := beego.AppConfig.String("httpport")
	_, err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		beego.Run()
	}
}
