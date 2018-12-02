package controllers

import (
	"github.com/louisevanderlith/mango/app/admin/logic"
	"github.com/louisevanderlith/mango/pkg/control"
)

type DefaultController struct {
	control.UIController
}

func NewDefaultCtrl(ctrlMap *control.ControllerMap) *DefaultController {
	result := &DefaultController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *DefaultController) Get() {
	c.Setup("default")
	c.CreateSideMenu(logic.GetMenu("/"))
}
