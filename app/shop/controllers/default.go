package controllers

import (
	"github.com/louisevanderlith/mango/util"
)

type DefaultController struct {
	util.UIController
}

func (c *DefaultController) Get() {
	c.Setup("main")
}
