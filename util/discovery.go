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

func GetServiceURL(instanceKey, serviceName string, cleanURL bool) (string, error) {
	cacheService, ok := serviceKeys[k{serviceName, cleanURL}]

	if !ok {
		contents, err := GETMessage("Router.API", "discovery", instanceKey, serviceName, strconv.FormatBool(cleanURL))

		if err != nil {
			return "", err
		}

		data := MarshalToResult(contents)

		if data.Failed {
			return "", data
		}

		result := data.Data.(string)
		serviceKeys[k{serviceName, cleanURL}] = result

		return result, nil
	}

	return cacheService, nil
}
