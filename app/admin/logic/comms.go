package logic

import (
	"github.com/louisevanderlith/mango/util"
)

type CommsObject struct {
	ID    int64
	Name  string
	Email string
	Phone string
	Body  string
}

func GetCommsMessages(instanceID string) ([]CommsObject, error) {
	var result []CommsObject
	contents, err := util.GETMessage(instanceID, "Communication.API", "message")

	if err != nil {
		return result, err
	}

	data := util.MarshalToResult(contents)

	if data.Failed() {
		return result, data
	}

	result = data.Data.([]CommsObject)

	return result, nil
}
