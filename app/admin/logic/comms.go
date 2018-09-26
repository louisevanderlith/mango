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
	resp := util.GETMessage(instanceID, "Communication.API", "message")

	if resp.Failed() {
		return []CommsObject{}, resp
	}

	result := resp.Data.([]CommsObject)

	return result, nil
}
