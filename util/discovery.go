package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"strconv"

	"github.com/louisevanderlith/mango/util/enums"
)

type Service struct {
	ID            string
	Name          string
	URL           string
	Version       int
	Environment   enums.Environment
	AllowedCaller enums.ServiceType
	Type          enums.ServiceType
}

type k struct {
	Name  string
	Clean bool
}

var (
	publicIP    string
	instanceKey string
	serviceKeys map[k]string
)

func init() {
	serviceKeys = make(map[k]string)

	serviceKeys[k{"Router.API", false}] = "http://localhost:8080/"
}

func GetInstanceKey() string {
	return instanceKey
}

// Register is used to register an application with the router service
func (service Service) Register(port string) (string, error) {
	service.URL = getPublicIP(port, service.Environment)

	contents, err := POSTMessage("Router.API", "discovery", service)

	if err != nil {
		return "", err
	}

	data := MarshalToResult(contents)

	if len(data.Error) != 0 {
		return "", errors.New(data.Error)
	}

	instanceKey = data.Data.(string)

	return instanceKey, nil
}

func GetServiceURL(serviceName string, cleanURL bool) (string, error) {
	var result string
	var finalError error

	cacheService, ok := serviceKeys[k{serviceName, cleanURL}]

	if ok {
		result = cacheService
	} else {
		contents, err := GETMessage("Router.API", "discovery", instanceKey, serviceName, strconv.FormatBool(cleanURL))

		if err != nil {
			return "", err
		}

		data := MarshalToResult(contents)

		if len(data.Error) != 0 {
			return "", errors.New(data.Error)
		}

		result = data.Data.(string)
		serviceKeys[k{serviceName, cleanURL}] = result
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
			log.Print("getPublicIP: ", err)
		}

		defer resp.Body.Close()

		ip, err := ioutil.ReadAll(resp.Body)
		publicIP = strings.Replace(string(ip), "\n", "", -1)

		if err != nil {
			log.Print("getPublicIP: ", err)
		}
	}

	return fmt.Sprintf("https://%s:%s/", publicIP, port)
}
