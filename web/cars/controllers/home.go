package controllers

import (
	"github.com/louisevanderlith/mango/pkg/control"
)

type HomeController struct {
	control.UIController
}

func NewHomeCtrl(ctrlMap *control.ControllerMap) *HomeController {
	result := &HomeController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *HomeController) Get() {
	c.Setup("home")
}
