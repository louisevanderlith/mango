package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/louisevanderlith/mango/util/enums"
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
	_publicIP    string
	_instanceKey string
	_serviceKeys map[string]string
)

func init() {
	_serviceKeys = make(map[string]string)
}

// Register is used to register an application with the router service
func (service Service) Register(discoveryURL string, port string) (string, error) {
	_serviceKeys["Router.API"] = discoveryURL

	service.URL = getPublicIP(port, service.Environment)

	contents, _ := POSTMessage("Router.API", "discovery", service)

	var data struct{ AppID string }
	jerr := json.Unmarshal(contents, &data)

	if jerr != nil {
		log.Print(jerr)
	}

	_instanceKey = data.AppID

	return _instanceKey, jerr
}

func GetServiceURL(serviceName string) (string, error) {
	var result string
	var finalError error

	cacheService, ok := _serviceKeys[serviceName]

	if ok {
		result = cacheService
	} else {

		contents, statusCode := GETMessage("Router.API", "discovery", _instanceKey, serviceName)

		var rawURL string
		err := json.Unmarshal(contents, &rawURL)

		if err != nil {
			log.Print(err)
		}

		if statusCode != 200 {
			finalError = errors.New(rawURL)
		} else {
			result = rawURL
			_serviceKeys[serviceName] = rawURL
		}
	}

	return result, finalError
}

func GETMessage(serviceName string, controller string, params ...string) ([]byte, int) {
	var result []byte
	var statusCode int
	url, err := GetServiceURL(serviceName)

	if err == nil {
		fullURL := fmt.Sprintf("%sv1/%s/%s", url, controller, strings.Join(params, "/"))
		resp, err := http.Get(fullURL)

		if err != nil {
			log.Print(err)
		} else {
			defer resp.Body.Close()
			statusCode = resp.StatusCode
			contents, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Print(err)
			}

			result = contents
		}
	} else {
		log.Print(err)
	}

	return result, statusCode
}

func POSTMessage(serviceName string, action string, obj interface{}) ([]byte, int) {
	var result []byte
	var statusCode int
	url, err := GetServiceURL(serviceName)

	if err == nil {
		fullURL := fmt.Sprintf("%sv1/%s", url, action)

		buff := new(bytes.Buffer)
		json.NewEncoder(buff).Encode(obj)

		resp, err := http.Post(fullURL, "application/json", buff)

		if err != nil {
			log.Print(err)
		} else {
			defer resp.Body.Close()
			statusCode = resp.StatusCode
			contents, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Print(err)
			}

			result = contents
		}
	}

	return result, statusCode
}

func getPublicIP(port string, env enums.Environment) string {

	if env == enums.DEV {
		_publicIP = "localhost"
	}

	if _publicIP == "" {
		resp, err := http.Get("http://myexternalip.com/raw")

		if err != nil {
			log.Print(err)
		}

		defer resp.Body.Close()

		ip, err := ioutil.ReadAll(resp.Body)
		_publicIP = strings.Replace(string(ip), "\n", "", -1)

		if err != nil {
			log.Print(err)
		}
	}

	return fmt.Sprintf("http://%s:%s/", _publicIP, port)
}
