package controllers

import (
	"github.com/louisevanderlith/mango/db/things"
	"encoding/json"
	"github.com/louisevanderlith/mango/util/control"
)

type SubCategoryController struct {
	control.APIController
}

// @Title GetSubCategory
// @Description Gets all Sub-Categories
// @Success 200 {string} string
// @router / [get]
func (req *SubCategoryController) Get() {
	var results []*things.Subcategory
	scat := things.Subcategory{}
	err := things.Ctx.SubCategory.Read(&scat, &results)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": results}
	}

	req.ServeJSON()
}

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
