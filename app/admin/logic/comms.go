package logic

import (
	"github.com/louisevanderlith/mango/util"
	"encoding/json"
	"log"
	"errors"
)

type CommsObject struct {
	Name  string
	Email string
	Phone string
	Body  string
}

func GetCommsMessages() ([]CommsObject, error) {
	var result []CommsObject
	var finalError error
	contents, statusCode := util.GETMessage("Communication.API", "message")

	if statusCode == 0 {
		if statusCode != 200 {
			var dataErr string
			err := json.Unmarshal(contents, &dataErr)

			if err != nil {
				log.Print(err)
			}

			finalError = errors.New(dataErr)
		} else {
			err := json.Unmarshal(contents, &result)

			if err != nil {
				log.Print(err)
			}
		}
	}

	return result, finalError
}
