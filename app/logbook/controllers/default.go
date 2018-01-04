package controllers

import (
	"github.com/louisevanderlith/mango/util/control"
)

type DefaultController struct {
	control.UIController
}

func (c *DefaultController) Get() {
	c.Setup("default")
}
