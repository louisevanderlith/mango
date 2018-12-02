package controllers

import (
	"github.com/louisevanderlith/mango/app/admin/logic"
	"github.com/louisevanderlith/mango/pkg/control"
)

type ManufacturerController struct {
	control.UIController
}

func NewManufacturerCtrl(ctrlMap *control.ControllerMap) *ManufacturerController {
	result := &ManufacturerController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *ManufacturerController) Get() {
	c.Setup("manufacturer")
	c.CreateSideMenu(logic.GetMenu("/manufacturer"))

	data, err := logic.GetManufacturers(c.GetInstanceID())

	c.Serve(data, err)
}
