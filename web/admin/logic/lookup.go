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
	resp, err := util.GETMessage(instanceID, "Things.API", "category")

	return toDTO(resp, err)
}

func GetManufacturers(instanceID string) ([]LookupObj, error) {
	resp, err := util.GETMessage(instanceID, "Things.API", "message")

	return toDTO(resp, err)
}

func GetModels(instanceID string) ([]LookupObj, error) {
	resp, err := util.GETMessage(instanceID, "Things.API", "model")

	return toDTO(resp, err)
}

func GetSubCategories(instanceID string) ([]LookupObj, error) {
	resp, err := util.GETMessage(instanceID, "Things.API", "subcategory")

	return toDTO(resp, err)
}

func toDTO(resp *util.RESTResult, err error) ([]LookupObj, error) {
	var result []LookupObj

	if err != nil {
		return result, err
	}

	if resp.Failed() {
		return result, resp
	}

	result, ok := resp.Data.([]LookupObj)

	if !ok {
		return result, errors.New("not a []LookupObj")
	}

	return result, nil
}
