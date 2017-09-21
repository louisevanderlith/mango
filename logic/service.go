package logic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

var (
	serviceMap map[string]Services
)

type Services []*Service

type Service struct {
	ID            string
	Name          string
	URL           string
	Version       int
	Requests      int
	Environment   string
	AllowedCaller string
	Type          string
}

func init() {
	serviceMap = make(map[string]Services)
	registerDatabases()
}

// Register is used to register an application with the router service
func Register(service Service) (string, error) {
	result := ""
	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(service)

	resp, err := http.Post(config.Discovery, "application/json", buff)

	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		var data struct{ AppID string }
		err := json.Unmarshal(contents, &data)

		if err != nil {
			log.Fatal(err)
		}

		result = data.AppID
	}

	return result, err
}

// AddService registers a new service and returns a key for that entry
func AddService(service Service) string {
	var result string
	items := serviceMap[service.Name]
	duplicate := false

	for _, value := range items {
		if value.URL == service.URL && value.Environment == service.Environment {
			duplicate = true
			result = value.ID
			break
		}
	}

	if !duplicate {
		u4, _ := uuid.NewV4()

		service.ID = u4.String()
		service.Version = getVersion()
		service.AllowedCaller = getAllowedCaller(service.Type)
		serviceMap[service.Name] = append(items, &service)

		result = service.ID
	}

	return result
}

// GetServicePath will return the correct URL for a requested service.
func GetServicePath(serviceName string, appID string) (string, error) {
	var result string
	var err error
	requestingApp := getRequestingService(appID)

	if requestingApp != nil {
		service := getService(serviceName, requestingApp.Environment, requestingApp.Type)

		if service != nil {
			result = service.URL
			service.Requests++
		} else {
			msg := fmt.Sprintf("%s wasn't found for the requesting application", serviceName)
			err = errors.New(msg)
		}
	} else {
		err = errors.New("Couldn't find an application with the given appID")
	}

	return result, err
}

func getAllowedCaller(serviceType string) string {
	var result string

	switch serviceType {
	case "database":
		result = "service"
	case "service":
		result = "proxy"
	case "proxy":
		result = "application"
	case "application":
		result = "*"
	}

	return result
}

func getService(serviceName string, environment string, callerType string) *Service {
	var result *Service
	serviceItems := serviceMap[serviceName]

	if serviceItems != nil {
		for _, val := range serviceItems {
			correctEnv := val.Environment == environment
			isAllowed := val.AllowedCaller == "*" || val.AllowedCaller == callerType

			if correctEnv && isAllowed {
				result = val
				break
			}
		}
	}

	return result
}

func getRequestingService(appID string) *Service {
	var result *Service

	for _, serviceItems := range serviceMap {
		for _, val := range serviceItems {
			if val.ID == appID {
				result = val
				break
			}
		}
	}

	return result
}

func registerDatabases() {
	for _, val := range *settings {
		for _, envVal := range val.Environments {
			db := Service{
				Environment: strings.ToLower(envVal.Name),
				Name:        val.Name,
				URL:         envVal.Value,
				Type:        "database"}

			AddService(db)
		}
	}
}

func getVersion() int {
	now := time.Now()
	concatDate := now.Format("0612")
	result, _ := strconv.Atoi(concatDate)

	return result
}
