package main

import (
	"log"

	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

var (
	config *util.Config
)

func init() {
	config = new(util.Config)
	config.LoadConfig("./app.conf")
}

func main() {
	// Register with router
	srv := util.Service{
		Environment: enums.GetEnvironment(config.Environment),
		Name:        "Proxy.API",
		URL:         config.Host,
		Type:        enums.PROXY}

	key, err := util.Register(srv, config.Discovery)

	if err != nil {
		log.Panic(err)
	} else {
		config.Key = key
		Proxer()
	}
}
