package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/db/comms"
	"github.com/louisevanderlith/mango/util/control"
)

type MessageController struct {
	control.APIController
}

// @Title SendMessage
// @Description Sends a Message
// @Param	body		body 	logic.Message	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *MessageController) Post() {
	var message comms.Message
	json.Unmarshal(req.Ctx.Input.RequestBody, &message)

	if message.To == "" {
		message.To = beego.AppConfig.String("defaultEmail")
	}

	err := message.SendMessage()

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "Message has been sent."}
	}

	req.ServeJSON()
}

// @Title GetMessages
// @Description Gets all Messages
// @Success 200 {string} string
// @router / [get]
func (req *MessageController) Get() {

	if req.Ctx.Output.Status != 401 {
		var result []*comms.Message
		msg := comms.Message{}
		err := comms.Ctx.Message.Read(msg, &result)

		if err != nil {
			req.Ctx.Output.SetStatus(500)
			req.Data["json"] = map[string]string{"Error": err.Error()}
		} else {
			req.Data["json"] = map[string]interface{}{"Data": result}
		}
	}

	req.ServeJSON()
}
