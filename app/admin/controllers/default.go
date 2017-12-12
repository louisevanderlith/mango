package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

type DefaultController struct {
	util.UIController
}

func init() {
	auths := make(util.ActionAuth)
	auths["GET"] = enums.Admin

	util.ProtectMethods(auths)
}

func (c *DefaultController) Get() {
	c.Setup("default")
}
