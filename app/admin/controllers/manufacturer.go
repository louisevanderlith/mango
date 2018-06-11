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

	c.Serve(err, data)
}
