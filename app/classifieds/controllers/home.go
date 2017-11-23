package controllers

import "github.com/louisevanderlith/mango/util"
import "github.com/louisevanderlith/mango/util/enums"

type HomeController struct {
	util.UIController
}

func init() {
	auths := make(map[string]enums.RoleType)
	auths["GET"] = enums.User

	util.ProtectMethods(auths)
}

func (c *HomeController) Get() {
	c.Setup("home")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
}
