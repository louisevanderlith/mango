package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/db/things"
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
	var results []*things.Category
	cat := things.Category{}
	err := things.Ctx.Category.Read(&cat, &results)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": results}
	}

	req.ServeJSON()
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

	_, err := things.Ctx.Category.Create(&obj)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "Save Successful."}
	}

	req.ServeJSON()
}
