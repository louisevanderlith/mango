package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/core/things"
	"github.com/louisevanderlith/mango/util/control"
)

type CategoryController struct {
	control.APIController
}

// @Title GetCategory
// @Description Gets all Categories
// @Success 200 {[]things.Category} []things.Category]
// @router / [get]
func (req *CategoryController) Get() {
	var results things.Categories
	cat := things.Category{}
	err := things.Ctx.Categories.Read(&cat, &results)

	req.Serve(err, results)
}

// @Title SaveCategory
// @Description Saves a new category
// @Param	body		body 	things.Category	true		"body for category"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *CategoryController) Post() {
	var obj things.Category
	json.Unmarshal(req.Ctx.Input.RequestBody, &obj)

	_, err := things.Ctx.Categories.Create(&obj)

	req.Serve(err, "Save Successful.")
}
