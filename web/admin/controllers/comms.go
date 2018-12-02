package controllers

import (
	"github.com/louisevanderlith/mango/app/admin/logic"
	"github.com/louisevanderlith/mango/pkg/control"
)

type CommsController struct {
	control.UIController
}

func NewCommsCtrl(ctrlMap *control.ControllerMap) *CommsController {
	result := &CommsController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *CommsController) Get() {
	c.Setup("comms")
	c.CreateSideMenu(logic.GetMenu("/comms"))

	data, err := logic.GetCommsMessages(c.GetInstanceID())

	c.Serve(data, err)
}
