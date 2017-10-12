package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/api/secure/logic"
)

type LoginController struct {
	util.BaseController
}

// @Title Login
// @Description Gets the form a user must fill in to login
// @Success 200 {string} string
// @router / [get]
func (req *LoginController) Get() {
	req.TplName = "login.html"
}

// @Title Login
// @Description Attempts to login against the provided credentials
// @Param	body		body 	logic.Login	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *LoginController) Post() {
	// failed logins should redirect to the login page
	var login logic.Login
	json.Unmarshal(req.Ctx.Input.RequestBody, &login)

	token := logic.AttemptLogin(login)

	if token == "" {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = "Login Failed"
	} else {
		req.Data["json"] = token
	}

	req.ServeJSON()
}

// @Title Logout
// @Description Logs out current logged in user session
// @Param	token		header 	string	true		"The session token"
// @Success 200 {string} logout success
// @router /logout [get]
func (req *LoginController) Logout() {
	token := req.GetString("token")

	if len(token) == 16 {
		logic.Delete(token)
		req.Data["json"] = "Logout Success"
	} else {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = "Invalid Token"
	}

	req.ServeJSON()
}
