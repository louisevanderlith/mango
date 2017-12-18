package main

import (
	_ "github.com/louisevanderlith/mango/app/www/routers"
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/util/enums"
	"log"
	"github.com/louisevanderlith/mango/util"
)

func main() {
	// Register with router
	srv := util.Service{
		Environment: enums.GetEnvironment(beego.BConfig.RunMode),
		Name:        beego.BConfig.AppName,
		Type:        enums.APP}

	port := beego.AppConfig.String("httpport")
	_, err := srv.Register(port)

	if err != nil {
		log.Printf("Register: ", err)
	} else {
		beego.Run()
	}
}

