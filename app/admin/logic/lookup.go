package logic

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/louisevanderlith/mango/util"
)

type LookupObj struct {
	ID          int64
	Name        string
	Description string
}

func GetCategories() ([]LookupObj, error) {
	contents, statusCode := util.GETMessage("Things.API", "category")

	return toDTO(contents, statusCode)
}

func GetManufacturers() ([]LookupObj, error) {
	contents, statusCode := util.GETMessage("Things.API", "message")

	return toDTO(contents, statusCode)
}

func GetModels() ([]LookupObj, error) {
	contents, statusCode := util.GETMessage("Things.API", "model")

	return toDTO(contents, statusCode)
}

func GetSubCategories() ([]LookupObj, error) {
	contents, statusCode := util.GETMessage("Things.API", "subcategory")

	return toDTO(contents, statusCode)
}

func toDTO(contents []byte, statusCode int) ([]LookupObj, error) {
	var result []LookupObj
	var finalErr error

	data := util.MarshalToMap(contents)

	if statusCode != 200 {
		var dataErr string
		err := json.Unmarshal(*data["Error"], &dataErr)

		if err != nil {
			log.Print("toDTO: ", err)
		}

		finalErr = errors.New(dataErr)
	} else {
		err := json.Unmarshal(*data["Data"], &result)

		if err != nil {
			log.Print("toDTO: ", err)
		}
	}

	return result, finalErr
}
