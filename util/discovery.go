package util

import (
	"errors"
	"fmt"
	"log"
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
	log.Printf("GetServiceURL:\tI:%s\tName:%s\n", instanceID, serviceName)
	cacheService, ok := serviceKeys[k{serviceName, cleanURL}]

	if !ok {
		contents, err := GETMessage(instanceID, "Router.API", "discovery", instanceID, serviceName, strconv.FormatBool(cleanURL))

		if err != nil {
			return "", err
		}

		data := MarshalToResult(contents)

		if data.Failed() {
			return "", data
		}

		if data.Data == nil {
			msg := fmt.Sprintf("data.Data is nil: %+v\n", data)
			return "", errors.New(msg)
		}

		result := data.Data.(string)
		serviceKeys[k{serviceName, cleanURL}] = result

		return result, nil
	}

	return cacheService, nil
}
