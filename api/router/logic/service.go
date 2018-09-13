package logic

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"

	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"

	"strings"

	uuid "github.com/nu7hatch/gouuid"
)

type Services []*util.Service

var serviceMap map[string]Services

func init() {
	serviceMap = make(map[string]Services)
}

// AddService registers a new service and returns a key for that entry
func AddService(service util.Service) (result string, err error) {
	items, ok := serviceMap[service.Name]

	if !ok {
		return "", errors.New(fmt.Sprintf("service.Name %s not found", service.Name))
	}

	duplicate := false

	for _, value := range items {
		if value.URL == service.URL && value.Environment == service.Environment {
			duplicate = true
			result = value.ID
			break
		}
	}

	if !duplicate {
		u4, err := uuid.NewV4()

		if err != nil {
			return "", err
		}

		service.ID = u4.String()
		service.Version = getVersion()
		service.AllowedCaller = getAllowedCaller(service.Type)
		serviceMap[service.Name] = append(items, &service)

		result = service.ID
	}

	return result, err
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

			result = "https://" + strings.ToLower(keyName) + cleanHost
		}
	} else {
		err = errors.New("Couldn't find an application with the given appID")
	}

	return result, err
}

func getCleanHost(env enums.Environment) string {
	envHost := fmt.Sprintf("HOST_%s", env)

	if len(envHost) == 0 {
		envHost = "HOST_DEV"
	}

	return beego.AppConfig.String(envHost)
}

func getAllowedCaller(serviceType enums.ServiceType) enums.ServiceType {
	var result enums.ServiceType

	switch serviceType {
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

func getVersion() int {
	now := time.Now()
	concatDate := now.Format("0612")
	result, _ := strconv.Atoi(concatDate)

	return result
}
