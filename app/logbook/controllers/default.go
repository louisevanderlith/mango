package controllers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

type MainController struct {
	util.UIController
}

func init(){
	auths := make(util.ActionAuth)
	auths["GET"] = enums.User

	util.ProtectMethods(auths)
}

func (c *MainController) Get() {
	c.Setup("main")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
