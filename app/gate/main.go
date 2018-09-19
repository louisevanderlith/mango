package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/astaxie/beego"

	"fmt"

	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
	"golang.org/x/net/http2"
)

func main() {
	// Register with router
	srv := util.Service{
		Environment: enums.GetEnvironment(beego.AppConfig.String("runmode")),
		Name:        beego.AppConfig.String("appname"),
		Type:        enums.APP}

	httpsPort := beego.AppConfig.String("httpsport")

	_, err := srv.Register(httpsPort)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		httpPort := beego.AppConfig.String("httpport")
		setupHost(httpPort, httpsPort)
	}
}

func setupHost(httpPort, httpsPort string) {
	registerSubdomains()

	//serveTLS(httpsPort)
	go serveHTTP2(httpsPort)

	err := http.ListenAndServe(":"+httpPort, http.HandlerFunc(redirectTLS))

	if err != nil {
		panic(err)
	}
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	moveURL := fmt.Sprintf("https://%s%s", r.Host, r.RequestURI)
	log.Printf("redirect: %s\n", moveURL)
	http.Redirect(w, r, moveURL, http.StatusPermanentRedirect)
}

func serveTLS(httpsPort string) {
	var srv http.Server
	srv.Addr = ":" + httpsPort
	srv.Handler = subdomains

	cerr := http2.ConfigureServer(&srv, nil)

	if cerr != nil {
		log.Print("ConfigureServer: ", cerr)
	}

	log.Println("Listening...")

	certPath := beego.AppConfig.String("certpath")
	hostCert := certPath + beego.AppConfig.String("hostCert")
	hostKey := certPath + beego.AppConfig.String("hostKey")

	err := srv.ListenAndServeTLS(hostCert, hostKey)

	if err != nil {
		panic(err)
	}
}

func serveHTTP2(httpsPort string) {
	certPath := beego.AppConfig.String("certpath")
	certPem := readCertBlock(certPath)
	keyPem := readKeyBlock(certPath)
	cert, err := tls.X509KeyPair(certPem, keyPem)

	if err != nil {
		panic(err)
	}

	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}

	srv := &http.Server{
		TLSConfig:    cfg,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		Addr:         ":" + httpsPort,
		Handler:      domains,
	}

	err = http2.ConfigureServer(srv, nil)

	if err != nil {
		panic(err)
	}

	log.Println("Listening...")

	err = srv.ListenAndServeTLS("", "")

	if err != nil {
		panic(err)
	}
}

func readBlocks(filePath string) []byte {
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	return file
}

func readCertBlock(path string) []byte {
	hostCert := path + beego.AppConfig.String("hostCert")

	return readBlocks(hostCert)
}

func readKeyBlock(path string) []byte {
	hostKey := path + beego.AppConfig.String("hostKey")

	return readBlocks(hostKey)
}
