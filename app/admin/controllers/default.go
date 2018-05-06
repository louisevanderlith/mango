package controllers

import "github.com/louisevanderlith/mango/app/admin/logic"

type DefaultController struct {
	logic.MenuController
}

func (c *DefaultController) Get() {
	c.Setup("default")
}
