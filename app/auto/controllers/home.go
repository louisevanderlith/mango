package controllers

import (
	"github.com/louisevanderlith/mango/util/control"
)

type HomeController struct {
	control.UIController
}

func (c *HomeController) Get() {
	c.Setup("home")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
}
