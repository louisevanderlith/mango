package mango

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type k struct {
	Name  string
	Clean bool
}

var serviceKeys map[k]string

func init() {
	serviceKeys = make(map[k]string)

	serviceKeys[k{"Router.API", false}] = fmt.Sprintf("http://Router%s:8080/", os.Getenv("RUNMODE"))
}

func GetServiceURL(instanceID, serviceName string, cleanURL bool) (string, error) {
	cacheService, ok := serviceKeys[k{serviceName, cleanURL}]
	log.Printf("[%t] Inst:\t%s\tService:%s\tClean:%t\n", ok, instanceID, serviceName, cleanURL)

	if !ok {
		resp, err := GETMessage(instanceID, "", "Router.API", "discovery", instanceID, serviceName, strconv.FormatBool(cleanURL))

		if err != nil {
			return "", err
		}

		if resp.Failed() {
			return "", resp
		}

		result := resp.Data.(string)
		serviceKeys[k{serviceName, cleanURL}] = result

		return result, nil
	}

	return cacheService, nil
}
