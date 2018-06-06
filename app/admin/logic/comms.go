package logic

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/louisevanderlith/mango/util"
)

type CommsObject struct {
	ID    int64
	Name  string
	Email string
	Phone string
	Body  string
}

func GetCommsMessages() ([]CommsObject, error) {
	var result []CommsObject
	var finalError error
	contents, statusCode := util.GETMessage("Communication.API", "message")
	data := util.MarshalToMap(contents)

	if statusCode != 200 {
		var dataErr string
		err := json.Unmarshal(*data["Error"], &dataErr)

		if err != nil {
			log.Print("GetCommsMessages: ", err)
		}

		finalError = errors.New(dataErr)
	} else {
		err := json.Unmarshal(*data["Data"], &result)

		if err != nil {
			log.Print("GetCommsMessages: ", err)
		}
	}

	return result, finalError
}
