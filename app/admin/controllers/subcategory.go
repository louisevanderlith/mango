package controllers

import (
	"github.com/louisevanderlith/mango/app/admin/logic"
)

type SubCategoryController struct {
	logic.MenuController
}

func (c *SubCategoryController) Get() {
	c.Setup("subcategory")

	data, err := logic.GetSubCategories()

	c.Serve(err, data)
}
