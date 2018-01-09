package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/louisevanderlith/mango/api/secure/logic"
	"github.com/louisevanderlith/mango/util/control"
)

type RegisterController struct {
	control.UIController
}

// @Title GetRegisterPage
// @Description Gets the form a user must fill in to register
// @Success 200 {string} string
// @router / [get]
func (req *RegisterController) Get() {
	req.Setup("register")
}

// @Title Register
// @Description Registers a new user
// @Param	body		body 	logic.Registration		true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *RegisterController) Post() {
	var user logic.Registration
	json.Unmarshal(req.Ctx.Input.RequestBody, &user)

	err := logic.SaveRegistration(user)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = err.Error()
	} else {
		req.Data["json"] = fmt.Sprintf("User %s created Successfully.", user.Name)
	}

	req.ServeJSON()
}
