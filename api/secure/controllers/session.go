package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/api/secure/logic"
)

type SessionController struct {
	util.SecureController
}

// @Title Get User Info
// @Description Gets the form a user must fill in to login
// @Success 200 {string} string
// @router / [get]
func (req *SessionController) Get() {
	// return the user's roles
	token := logic.GetAvoToken(req.Ctx)
	roles := logic.GetRoles(token)

	if len(roles) > 0 {
		req.Data["json"] = map[string]interface{}{"Data": roles}
	} else {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": "No Roles Found"}
	}

	req.ServeJSON()
}