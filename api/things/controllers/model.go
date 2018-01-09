package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/db/things"
	"github.com/louisevanderlith/mango/util/control"
)

type ModelController struct {
	control.APIController
}

// @Title GetModel
// @Description Gets all Models
// @Success 200 {[]things.Model} []things.Model
// @router / [get]
func (req *ModelController) Get() {
	var results []*things.Model
	mdl := things.Model{}
	err := things.Ctx.Model.Read(mdl, &results)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": results}
	}

	req.ServeJSON()
}

// @Title SaveModel
// @Description Saves a new model
// @Param	body		body 	things.Model	true		"body for model"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *ModelController) Post() {
	var obj things.Model
	json.Unmarshal(req.Ctx.Input.RequestBody, &obj)

	_, err := things.Ctx.Model.Create(&obj)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "Save Successful."}
	}

	req.ServeJSON()
}
