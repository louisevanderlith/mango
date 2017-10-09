package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/db/secure"
)

type RegisterController struct {
	beego.Controller
}

func (req *RegisterController) Get() {
	req.TplName = "register.html"
}

// @Title Register
// @Description Registers a new user
// @Param	body		body 	secure.User	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *RegisterController) Post() {
	var user secure.User
	json.Unmarshal(req.Ctx.Input.RequestBody, &user)

	err := secure.CreateUser(user)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = err.Error()
	} else {
		req.Data["json"] = fmt.Sprintf("User %s created Successfully.", user.Name)
	}

	req.ServeJSON()
}
