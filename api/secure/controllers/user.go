package controllers

import (
	"github.com/louisevanderlith/mango/api/secure/logic"
	"github.com/louisevanderlith/mango/util/control"
)

type UserController struct {
	control.APIController
}

// @Title GetUsers
// @Description Gets all Users
// @Success 200 {[]logic.UserObject]} []logic.UserObject]
// @router / [get]
func (req *UserController) Get() {
	if req.Ctx.Output.Status != 401 {
		result, err := logic.GetUsers()

		if err != nil {
			req.Ctx.Output.SetStatus(500)
			req.Data["json"] = map[string]string{"Error": err.Error()}
		} else {
			req.Data["json"] = map[string]interface{}{"Data": result}
		}
	}

	req.ServeJSON()
}
