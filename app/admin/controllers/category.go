package controllers

import (
	"github.com/louisevanderlith/mango/app/admin/logic"
	"github.com/louisevanderlith/mango/util/control"
)

type CategoryController struct {
	control.UIController
}

func (c *CategoryController) Get() {
	c.Setup("category")

	data, err := logic.GetCategories()

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["error"] = err
	} else {
		c.Data["data"] = data
	}
}
