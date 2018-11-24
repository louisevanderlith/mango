package util

import (
	"strconv"
)

type k struct {
	Name  string
	Clean bool
}

var serviceKeys map[k]string

func init() {
	serviceKeys = make(map[k]string)

	//this is hard coded for a reason, keep it that way.
	serviceKeys[k{"Router.API", false}] = "http://localhost:8080/"
}

func GetServiceURL(instanceID, serviceName string, cleanURL bool) (string, error) {
	cacheService, ok := serviceKeys[k{serviceName, cleanURL}]

	if !ok {
		resp, err := GETMessage(instanceID, "Router.API", "discovery", instanceID, serviceName, strconv.FormatBool(cleanURL))

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
