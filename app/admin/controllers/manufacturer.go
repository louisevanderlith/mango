package controllers

import (
	"github.com/louisevanderlith/mango/app/admin/logic"
	"github.com/louisevanderlith/mango/util/control"
)

type ManufacturerController struct {
	control.UIController
}

func (c *ManufacturerController) Get() {
	c.Setup("manufacturer")
	c.CreateSideMenu(logic.GetMenu("/manufacturer"))

	data, err := logic.GetManufacturers()

	c.Serve(err, data)
}
