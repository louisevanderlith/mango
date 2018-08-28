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
	c.CreateSideMenu(logic.GetMenu("/subcategory"))

	data, err := logic.GetSubCategories()

	c.Serve(err, data)
}
