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

	data, err := logic.GetManufacturers()

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["error"] = err
	} else {
		c.Data["data"] = data
	}
}
