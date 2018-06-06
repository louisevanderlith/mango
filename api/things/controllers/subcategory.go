package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/db/things"
	"github.com/louisevanderlith/mango/util/control"
)

type SubCategoryController struct {
	control.APIController
}

// @Title GetSubCategory
// @Description Gets all Sub-Categories
// @Success 200 {[]things.Subcategory} []things.Subcategory
// @router / [get]
func (req *SubCategoryController) Get() {
	var result things.Subcategories
	scat := things.Subcategory{}
	err := things.Ctx.SubCategories.Read(&scat, &result)

	req.Serve(err, result)
}

// @Title SaveSubcategory
// @Description Saves a new sub-category
// @Param	body		body 	things.Subcategory	true		"body for sub-category"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *SubCategoryController) Post() {
	var obj things.Subcategory
	json.Unmarshal(req.Ctx.Input.RequestBody, &obj)

	_, err := things.Ctx.SubCategories.Create(&obj)

	req.Serve(err, "Save Successful.")
}
