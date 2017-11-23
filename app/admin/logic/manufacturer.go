package logic

import (
	"encoding/json"
	"log"
	"github.com/louisevanderlith/mango/util"
	"errors"
)

type Manufacturer struct {
	ID   int64
	Name string
}

func GetManufacturers() ([]Manufacturer, error){
	var result []Manufacturer
	var finalError error
	contents, statusCode := util.GETMessage("Things.API", "message", "")
	data := util.MarshalToMap(contents)

	if statusCode != 200 {
		var dataErr string
		err := json.Unmarshal(*data["Error"], &dataErr)

		if err != nil {
			log.Printf("GetCommsMessages: ", err)
		}

		finalError = errors.New(dataErr)
	} else {
		err := json.Unmarshal(*data["Data"], &result)

		if err != nil {
			log.Printf("GetCommsMessages: ", err)
		}
	}

	return result, finalError
}
