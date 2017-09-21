package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/louisevanderlith/mango/api/comms/routers"
	"github.com/louisevanderlith/mango/logic"

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
	srv := logic.Service{
		Environment: beego.BConfig.RunMode,
		Name:        beego.BConfig.AppName,
		URL:         "http://localhost:" + beego.AppConfig.String("httpport"),
		Type:        "service"
	}

	key, err := logic.Register()

	if err != nil{
		log.Panic(err)
	} else {
		instanceKey = key
		logic.BuildDatabase(instanceKey)
		beego.Run()
	}
}
