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
	var result []*things.Subcategory
	scat := things.Subcategory{}
	err := things.Ctx.SubCategory.Read(&scat, &result)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": result}
	}

	req.ServeJSON()
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

	_, err := things.Ctx.SubCategory.Create(&obj)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "Save Successful."}
	}

	req.ServeJSON()
}
