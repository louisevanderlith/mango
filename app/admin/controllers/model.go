package controllers

import (
	"github.com/louisevanderlith/mango/app/admin/logic"
)

type ModelController struct {
	logic.MenuController
}

func (c *ModelController) Get() {
	c.Setup("model")

	data, err := logic.GetModels()

	c.Serve(err, data)
}
