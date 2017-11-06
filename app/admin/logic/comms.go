package logic

import "github.com/louisevanderlith/mango/util"

type CommsObject struct {
	Name  string
	Email string
	Phone string
	Body  string
}

func GetCommsMessages() {
	content, statusCode := util.GETMessage("Communication.API", "message")

}
