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
	c.CreateSideMenu(logic.GetMenu("/category"))

	data, err := logic.GetCategories()

	c.Serve(err, data)
}
