package controllers

import (
	"github.com/louisevanderlith/mango/app/admin/logic"
	"github.com/louisevanderlith/mango/util/control"
)

type ModelController struct {
	control.UIController
}

func NewModelCtrl(ctrlMap *control.ControllerMap) *ModelController {
	result := &ModelController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *ModelController) Get() {
	c.Setup("model")
	c.CreateSideMenu(logic.GetMenu("/model"))

	data, err := logic.GetModels(c.GetInstanceID())

	c.Serve(data, err)
}
