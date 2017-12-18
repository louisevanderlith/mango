package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/louisevanderlith/mango/util/enums"
	"github.com/astaxie/beego"
)

type Service struct {
	ID            string
	Name          string
	URL           string
	Version       int
	Requests      int
	Environment   enums.Environment
	AllowedCaller enums.ServiceType
	Type          enums.ServiceType
}

var (
	publicIP    string
	instanceKey string
	serviceKeys map[string]string
)

func init() {
	serviceKeys = make(map[string]string)

	if beego.AppConfig.String("runmode") == "dev" {
		serviceKeys["Router.API"] = "http://localhost:8080/"
	} else {
		serviceKeys["Router.API"] = "https://router.avosa.co.za/"
	}
}

func GetInstanceKey() string {
	return instanceKey
}

// Register is used to register an application with the router service
func (service Service) Register(port string) (string, error) {
	service.URL = getPublicIP(port, service.Environment)

	contents, _ := POSTMessage("Router.API", "discovery", service)

	var data struct{ AppID string }
	jerr := json.Unmarshal(contents, &data)

	if jerr != nil {
		log.Printf("json.Unmarshal: ", jerr)
	}

	instanceKey = data.AppID

	return instanceKey, jerr
}

func GetServiceURL(serviceName string) (string, error) {
	var result string
	var finalError error

	cacheService, ok := serviceKeys[serviceName]

	if ok {
		result = cacheService
	} else {
		contents, statusCode := GETMessage("Router.API", "discovery", instanceKey, serviceName)

		var rawURL string
		err := json.Unmarshal(contents, &rawURL)

		if err != nil {
			log.Printf("json.Unmarshal: ", err)
		}

		if statusCode != 200 {
			finalError = errors.New(rawURL)
		} else {
			result = rawURL
			serviceKeys[serviceName] = rawURL
		}
	}

	return result, finalError
}

func getPublicIP(port string, env enums.Environment) string {

	if env == enums.DEV {
		publicIP = "localhost"
	}

	if publicIP == "" {
		resp, err := http.Get("http://myexternalip.com/raw")

		if err != nil {
			log.Printf("getPublicIP: ", err)
		}

		defer resp.Body.Close()

		ip, err := ioutil.ReadAll(resp.Body)
		publicIP = strings.Replace(string(ip), "\n", "", -1)

		if err != nil {
			log.Printf("getPublicIP: ", err)
		}
	}

	return fmt.Sprintf("http://%s:%s/", publicIP, port)
}
