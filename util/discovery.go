package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

// Register is used to register an application with the router service
func Register(service Service, discoveryURL string) (string, error) {
	result := ""
	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(service)

	resp, err := http.Post(discoveryURL, "application/json", buff)

	if err != nil {
		log.Panic(err)
	} else {
		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Panic(err)
		}

		var data struct{ AppID string }
		jerr := json.Unmarshal(contents, &data)

		if jerr != nil {
			log.Panic(jerr)
			err = jerr
		}

		result = data.AppID
	}

	return result, err
}

func GetServiceURL(instanceKey string, serviceName string, discoveryURL string) (string, error) {
	var result string
	var err error

	discoveryRoute := fmt.Sprintf("%s%s/%s", discoveryURL, instanceKey, serviceName)
	resp, err := http.Get(discoveryRoute)
	defer resp.Body.Close()

	if err != nil {
		log.Panic(err)
	} else {
		contents, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Panic(err)
		}

		jsonErr := json.Unmarshal(contents, &result)

		if jsonErr != nil {
			log.Panic(err)
		}

		if result == "" {
			msg := fmt.Sprintf("Couldn't find a application for %s", serviceName)
			err = errors.New(msg)
		}
	}

	log.Print(result)
	return result, err
}
