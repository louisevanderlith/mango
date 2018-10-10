package controllers

import (
	"github.com/louisevanderlith/mango/core/secure"
	"github.com/louisevanderlith/mango/util/control"
)

type UserController struct {
	control.APIController
}

func NewUserCtrl(ctrlMap *control.ControllerMap) *UserController {
	result := &UserController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetUsers
// @Description Gets all Users
// @Success 200 {[]logic.UserObject]} []logic.UserObject]
// @router /:pageData [get]
func (req *UserController) Get() {
	page, size := req.GetPageData()
	result := secure.GetUsers(page, size)
	req.Serve(result, nil)
}
