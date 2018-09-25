package main

import (
	"log"

	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"

	"github.com/louisevanderlith/mango/api/comment/routers"
	_ "github.com/louisevanderlith/mango/api/comms/routers"
	_ "github.com/louisevanderlith/mango/core/comms"

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
	srv := util.NewService(mode, name, enums.API)

	log.Printf("%+v\n", srv)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		showSMTPInfo()
		beego.Run()
	}
}

func showSMTPInfo() {
	smtpUser := beego.AppConfig.String("smtpUsername")
	smtpAddress := beego.AppConfig.String("smtpAddress")
	smtpPort := beego.AppConfig.String("smtpPort")

	log.Print(smtpUser, smtpAddress, smtpPort)
}
