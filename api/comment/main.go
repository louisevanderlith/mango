package main

import (
	"log"

	_ "github.com/louisevanderlith/mango/api/comment/routers"
	"github.com/louisevanderlith/mango/db/comment"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

var instanceKey string

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Register with router
	srv := util.Service{
		Environment: enums.GetEnvironment(beego.BConfig.RunMode),
		Name:        beego.BConfig.AppName,
		Type:        enums.API}

	discURL := beego.AppConfig.String("discovery")
	port := beego.AppConfig.String("httpport")
	key, err := util.Register(srv, discURL, port)

	if err != nil {
		log.Print(err)
	} else {
		instanceKey = key
		comment.NewDatabase(instanceKey, discURL)
		beego.Run()
	}
}
