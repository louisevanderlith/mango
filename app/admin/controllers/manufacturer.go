package controllers

import (
	"github.com/louisevanderlith/mango/app/admin/logic"
)

type ManufacturerController struct {
	logic.MenuController
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
