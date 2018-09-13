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

func GetCategories() ([]LookupObj, error) {
	contents, err := util.GETMessage("Things.API", "category")

	if err != nil {
		return []LookupObj{}, err
	}

	return toDTO(contents)
}

func GetManufacturers() ([]LookupObj, error) {
	contents, err := util.GETMessage("Things.API", "message")

	if err != nil {
		return []LookupObj{}, err
	}

	return toDTO(contents)
}

func GetModels() ([]LookupObj, error) {
	contents, err := util.GETMessage("Things.API", "model")

	if err != nil {
		return []LookupObj{}, err
	}

	return toDTO(contents)
}

func GetSubCategories() ([]LookupObj, error) {
	contents, err := util.GETMessage("Things.API", "subcategory")

	if err != nil {
		return []LookupObj{}, err
	}

	return toDTO(contents)
}

func toDTO(contents []byte) ([]LookupObj, error) {
	var result []LookupObj

	data := util.MarshalToResult(contents)

	if len(data.Error) != 0 {
		return result, errors.New(data.Error)
	}

	result, ok := data.Data.([]LookupObj)

	if !ok {
		return result, errors.New("not a []LookupObj")
	}

	return result, nil
}
