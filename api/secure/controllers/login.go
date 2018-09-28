package controllers

import (
	"errors"

	"github.com/louisevanderlith/mango/api/secure/logic"
	"github.com/louisevanderlith/mango/util/control"
)

type LoginController struct {
	control.UIController
}

func NewLoginCtrl(ctrlMap *control.ControllerMap) *LoginController {
	result := &LoginController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetLoginPage
// @Description Gets the form a user must fill in to login
// @Success 200 {string} string
// @router / [get]
func (req *LoginController) Get() {
	req.Setup("login")
}

// @Title GetAvo
// @Description Gets the currently logged in user's avo
// @Param	path	path	string	true	"sessionID"
// @Success 200 {map[string]string} map[string]string
// @router /avo/:sessionID [get]
func (req *LoginController) GetAvo() {
	sessionID := req.Ctx.Input.Param(":sessionID")
	hasAvo := control.HasAvo(sessionID)

	var err error
	var result control.Cookies

	if !hasAvo {
		err = errors.New("no data found")
	} else {
		result = control.FindAvo(sessionID)
	}

	req.APIController.Serve(result, err)
}

// @Title Login
// @Description Attempts to login against the provided credentials
// @Param	body		body 	logic.Login	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *LoginController) Post() {
	sessionID, err := logic.AttemptLogin(req.Ctx)

	req.APIController.Serve(sessionID, err)
}

// @Title Logout
// @Description Logs out current logged in user session
// @Param	path	path	string	true	"sessionID"
// @Success 200 {string} string
// @router /logout/:sessionID [get]
func (req *LoginController) Logout() {
	sessionID := req.Ctx.Input.Param(":sessionID")

	// TODO: Create Trace for Logout...
	control.DestroyAvo(sessionID)

	req.APIController.Serve("Logout Success", nil)
}
