package controllers

import (
	"github.com/louisevanderlith/mango/core/secure"
	"github.com/louisevanderlith/mango/util/control"
)

type UserController struct {
	control.APIController
}

// @Title GetUsers
// @Description Gets all Users
// @Success 200 {[]logic.UserObject]} []logic.UserObject]
// @router /:pageData(^[A-Z](?:_?[0-9]+)*$) [get]
func (req *UserController) Get() {
	page, size := req.GetPageData()
	result, err := secure.GetUsers(page, size)
	req.Serve(err, result)
}
