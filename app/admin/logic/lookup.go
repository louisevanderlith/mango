package logic

import (
	"errors"

	"github.com/louisevanderlith/mango/util"
)

type LookupObj struct {
	ID          int64
	Name        string
	Description string
}

func GetCategories(instanceID string) ([]LookupObj, error) {
	contents, err := util.GETMessage(instanceID, "Things.API", "category")

	if err != nil {
		return []LookupObj{}, err
	}

	return toDTO(contents)
}

func GetManufacturers(instanceID string) ([]LookupObj, error) {
	contents, err := util.GETMessage(instanceID, "Things.API", "message")

	if err != nil {
		return []LookupObj{}, err
	}

	return toDTO(contents)
}

func GetModels(instanceID string) ([]LookupObj, error) {
	contents, err := util.GETMessage(instanceID, "Things.API", "model")

	if err != nil {
		return []LookupObj{}, err
	}

	return toDTO(contents)
}

func GetSubCategories(instanceID string) ([]LookupObj, error) {
	contents, err := util.GETMessage(instanceID, "Things.API", "subcategory")

	if err != nil {
		return []LookupObj{}, err
	}

	return toDTO(contents)
}

func toDTO(contents []byte) ([]LookupObj, error) {
	var result []LookupObj

	data := util.MarshalToResult(contents)

	if data.Failed() {
		return result, data
	}

	result, ok := data.Data.([]LookupObj)

	if !ok {
		return result, errors.New("not a []LookupObj")
	}

	return result, nil
}
