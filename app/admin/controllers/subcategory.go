package controllers

import (
	"github.com/louisevanderlith/mango/app/admin/logic"
	"github.com/louisevanderlith/mango/util/control"
)

type SubCategoryController struct {
	control.UIController
}

func (c *SubCategoryController) Get() {
	c.Setup("subcategory")

	data, err := logic.GetSubCategories()

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["error"] = err
	} else {
		c.Data["data"] = data
	}
}
