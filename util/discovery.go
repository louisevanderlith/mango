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

var publicIP string

// Register is used to register an application with the router service
func Register(service Service, discoveryURL string, port string) (string, error) {
	result := ""

	service.URL = getPublicIP(port, service.Environment)

	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(service)

	resp, err := http.Post(discoveryURL, "application/json", buff)

	if err != nil {
		log.Print(err)
	} else {
		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Print(err)
		}

		var data struct{ AppID string }
		jerr := json.Unmarshal(contents, &data)

		if jerr != nil {
			log.Print(jerr)
			err = jerr
		}

		result = data.AppID
	}

	return result, err
}

func GetServiceURL(instanceKey string, serviceName string, discoveryURL string) (string, error) {
	var result string
	var finalError error

	discoveryRoute := fmt.Sprintf("%s%s/%s", discoveryURL, instanceKey, serviceName)
	resp, err := http.Get(discoveryRoute)

	if err != nil {
		log.Print(err)
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Print(err)
		}

		var rawURL string
		err = json.Unmarshal(contents, &rawURL)

		if err != nil {
			log.Print(err)
		}

		if resp.StatusCode != 200 {
			finalError = errors.New(rawURL)
		} else {
			result = rawURL
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
			log.Print(err)
		}

		defer resp.Body.Close()

		ip, err := ioutil.ReadAll(resp.Body)
		publicIP = strings.Replace(string(ip), "\n", "", -1)

		if err != nil {
			log.Print(err)
		}
	}

	return fmt.Sprintf("http://%s:%s/", publicIP, port)
}
