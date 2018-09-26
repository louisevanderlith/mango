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
	resp := util.GETMessage(instanceID, "Things.API", "category")

	return toDTO(resp)
}

func GetManufacturers(instanceID string) ([]LookupObj, error) {
	resp := util.GETMessage(instanceID, "Things.API", "message")

	return toDTO(resp)
}

func GetModels(instanceID string) ([]LookupObj, error) {
	resp := util.GETMessage(instanceID, "Things.API", "model")

	return toDTO(resp)
}

func GetSubCategories(instanceID string) ([]LookupObj, error) {
	resp := util.GETMessage(instanceID, "Things.API", "subcategory")

	return toDTO(resp)
}

func toDTO(resp *util.RESTResult) ([]LookupObj, error) {
	var result []LookupObj

	if resp.Failed() {
		return result, resp
	}

	result, ok := resp.Data.([]LookupObj)

	if !ok {
		return result, errors.New("not a []LookupObj")
	}

	return result, nil
}
