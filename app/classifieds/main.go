package main

import (
	"log"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	_ "github.com/louisevanderlith/classifieds/routers"
	"github.com/louisevanderlith/mango/logic"
)

var instanceKey string

func main() {
	// Register with router
	srv := logic.Service{
		Environment: "dev",
		Name:        "Classifieds.APP",
		URL:         "http://localhost:xxx",
		Type:        "application"}

	key, err := logic.Register()

	if err != nil {
		log.Panic(err)
	} else {
		instanceKey = key
		beego.Run()
	}
}
