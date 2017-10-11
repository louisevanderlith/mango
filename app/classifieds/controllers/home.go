package controllers

import "github.com/louisevanderlith/mango/util"

type HomeController struct {
	util.BaseController
}

func (c *HomeController) Get() {
	c.Setup("home")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
}
