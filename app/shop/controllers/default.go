package controllers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/util"
)

type MainController struct {
	util.UIController
}

func (c *MainController) Get() {
	c.Setup("main")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
