package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/db/comms"
)

type MessageController struct {
	beego.Controller
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
	msg := comms.Message{}
	result, err := msg.ReadAll()

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": result}
	}

	req.ServeJSON()
}
