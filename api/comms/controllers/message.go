package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/core/comms"
	"github.com/louisevanderlith/mango/util/control"
)

type MessageController struct {
	control.APIController
}

// @Title SendMessage
// @Description Sends a Message
// @Param	body	body	comms.Message	true	"body for message content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *MessageController) Post() {
	var message comms.Message
	json.Unmarshal(req.Ctx.Input.RequestBody, &message)

	if message.To == "" {
		message.To = beego.AppConfig.String("defaultEmail")
	}

	err := message.SendMessage()

	req.Serve(err, "Message has been sent.")
}

// @Title GetMessages
// @Description Gets all Messages
// @Success 200 {[]comms.Message]} []comms.Message]
// @router /:pageData(^[A-Z](?:_?[0-9]+)*$) [get]
func (req *MessageController) Get() {
	page, size := req.GetPageData()
	result, err := comms.GetMessages(page, size)

	req.Serve(err, result)
}
