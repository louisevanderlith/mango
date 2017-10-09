package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/db/comms"
)

// Operations about Messages
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

	err := comms.SendMessage(message)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "Message has been sent."}
	}

	req.ServeJSON()
}
