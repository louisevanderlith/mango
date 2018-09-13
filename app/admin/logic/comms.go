package logic

import (
	"errors"

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
	contents, err := util.GETMessage("Communication.API", "message")

	if err != nil {
		return result, err
	}

	data := util.MarshalToResult(contents)

	if len(data.Error) != 0 {
		return result, errors.New(data.Error)
	}

	result = data.Data.([]CommsObject)

	return result, nil
}
