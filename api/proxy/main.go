package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/louisevanderlith/mango/logic"
)

var (
	config *logic.Config
)

func init() {
	config = new(logic.Config)
	config.LoadConfig("./app.conf")
}

func main() {
	// Register with router
	srv := logic.Service{
		Environment: config.Environment,
		Name:        "Proxy.API",
		URL:         config.Host,
		Type:        "proxy"
	}

	key, err := logic.Register()
	
	if err != nil {
		log.Panic(err)
	} else {
		config.Key = key
		Proxer()
	}
}
