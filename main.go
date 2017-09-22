package main

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/mango/logic"
)

var instanceKey string

func main() {
	// Register with router
	srv := logic.Service{
		Environment: "dev",
		Name:        "Website.APP",
		URL:         "http://localhost:80",
		Type:        "application"}

	discURL := "http://localhost:123"
	key, err := logic.Register(srv, discURL)

	if err != nil {
		log.Panic(err)
	} else {
		instanceKey = key
		setupHost()
	}
}

func setupHost() {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))
	http.Handle("/", fs)

	log.Println("Listening...")
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}
}
