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
	result, err := logic.GetUsers()

	req.Serve(err, result)
}
