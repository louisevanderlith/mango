package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/core/secure"
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
// @Param	body		body 	secure.AuthRequest		true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *RegisterController) Post() {
	var regis secure.Registration
	json.Unmarshal(req.Ctx.Input.RequestBody, &regis)

	result, err := secure.Register(regis)
	req.Serve(err, result)
}
