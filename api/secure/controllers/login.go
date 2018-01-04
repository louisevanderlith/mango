package controllers

import (
	"github.com/louisevanderlith/mango/api/secure/logic"
	"github.com/louisevanderlith/mango/util"
	"net/http"
	"strings"
	"github.com/louisevanderlith/mango/util/control"
)

type LoginController struct {
	control.UIController
}

// @Title GetLogin
// @Description Gets the form a user must fill in to login
// @Param	query	query	string	true	"return URL"
// @Success 200 {string} string
// @router / [get]
func (req *LoginController) Get() {
	sessionID := req.Ctx.GetCookie("beegosessionID")
	hasAvo := util.HasAvo(sessionID)

	if !hasAvo {
		req.Setup("login")
	} else {
		ret := req.Ctx.Input.Query("return")

		if ret != "" && strings.HasPrefix(ret, "http") {
			req.Redirect(ret, http.StatusTemporaryRedirect)
		}
	}
}

// @Title GetCookie
// @Description Gets the currently logged in user's cookie
// @Param	path	path	string	true	"return URL"
// @Success 200 {string} string
// @router /avo/:sessionID [get]
func (req *LoginController) GetCookie() {
	sessionID := req.Ctx.Input.Param(":sessionID")
	hasAvo := util.HasAvo(sessionID)

	if !hasAvo {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": "No data found."}
	} else {
		data := util.FindAvo(sessionID)
		req.Data["json"] = map[string]interface{}{"Data": data}
	}

	req.ServeJSON()
}

// @Title Login
// @Description Attempts to login against the provided credentials
// @Param	body		body 	logic.Login	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *LoginController) Post() {
	loggedIn, sessionID, err := logic.AttemptLogin(req.Ctx)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = "Login Error " + err.Error()
	} else if !loggedIn {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = "Login Failed"
	} else {
		req.Data["json"] = sessionID
	}

	req.ServeJSON()
}

// @Title Logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (req *LoginController) Logout() {
	sessionID := req.Ctx.GetCookie("beegosessionID")
	util.DestroyAvo(sessionID)

	req.Data["json"] = "Logout Success"
	req.ServeJSON()
}
