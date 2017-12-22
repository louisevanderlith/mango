package logic

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"

	uuid "github.com/nu7hatch/gouuid"
	"strings"
)

type Services []*util.Service

var (
	serviceMap map[string]Services
)

func init() {
	serviceMap = make(map[string]Services)
	registerDatabases()
}

// AddService registers a new service and returns a key for that entry
func AddService(service util.Service) string {
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
func GetServicePath(serviceName string, appID string, clean bool) (string, error) {
	var result string
	var err error
	requestingApp := getRequestingService(appID)

	if requestingApp != nil {
		if !clean {
			service := getService(serviceName, requestingApp.Environment, requestingApp.Type)

			if service != nil {
				result = service.URL
			} else {
				msg := fmt.Sprintf("%s wasn't found for the requesting application", serviceName)
				err = errors.New(msg)
			}
		} else {
			keyName := strings.Split(serviceName, ".")[0]
			cleanHost := getCleanHost(requestingApp.Environment)

			result = strings.ToLower(keyName) + cleanHost
		}
	} else {
		err = errors.New("Couldn't find an application with the given appID")
	}

	return result, err
}

func getCleanHost(env enums.Environment) string {
	var result string
	switch env {
	case enums.LIVE:
		result = ".avosa.co.za/"
	case enums.UAT:
		result = ".???.co.za/"
	case enums.DEV:
		result = ".localhost/"
	default:
		result = ".localhost/"
	}

	return result
}

func getAllowedCaller(serviceType enums.ServiceType) enums.ServiceType {
	var result enums.ServiceType

	switch serviceType {
	case enums.DB:
		result = enums.API
	case enums.API:
		result = enums.APP
	case enums.APP:
		result = enums.ANY
	}

	return result
}

func getService(serviceName string, environment enums.Environment, callerType enums.ServiceType) *util.Service {
	var result *util.Service
	serviceItems := serviceMap[serviceName]

	if serviceItems != nil {
		for _, val := range serviceItems {
			correctEnv := val.Environment == environment
			isAllowed := val.AllowedCaller == enums.ANY || val.AllowedCaller == callerType

			if correctEnv && isAllowed {
				result = val
				break
			}
		}
	}

	return result
}

func getRequestingService(appID string) *util.Service {
	var result *util.Service

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
	settings := loadSettings()

	for _, val := range *settings {
		for _, envVal := range val.Environments {
			db := util.Service{
				Environment: enums.GetEnvironment(envVal.Name),
				Name:        val.Name,
				URL:         envVal.Value,
				Type:        enums.DB}

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
