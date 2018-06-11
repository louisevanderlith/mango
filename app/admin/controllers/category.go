package controllers

import (
	"github.com/louisevanderlith/mango/app/admin/logic"
)

type CategoryController struct {
	logic.MenuController
}

func (c *CategoryController) Get() {
	c.Setup("category")

	data, err := logic.GetCategories()

	c.Serve(err, data)
}
