package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

type MainController struct {
	util.UIController
}

func init() {
	auths := make(util.ActionAuth)
	auths["GET"] = enums.Admin

	util.ProtectMethods(auths)
}

func (c *MainController) Get() {
	c.Setup("main")
}
