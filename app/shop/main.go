package main

import (
	"log"

	"github.com/astaxie/beego"
	_ "github.com/louisevanderlith/mango/app/shop/routers"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

func main() {
	// Register with router
	srv := util.Service{
		Environment: enums.GetEnvironment(beego.BConfig.RunMode),
		Name:        beego.BConfig.AppName,
		Type:        enums.APP}

	port := beego.AppConfig.String("httpport")
	_, err := srv.Register( port)

	if err != nil {
		log.Print(err)
	} else {
		beego.Run()
	}
}
