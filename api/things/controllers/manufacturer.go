package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/db/things"
	"github.com/louisevanderlith/mango/util/control"
)

type ManufacturerController struct {
	control.APIController
}

// @Title GetManufacturer
// @Description Gets all Manufacturers
// @Success 200 {[]things.Manufacturer} []things.Manufacturer
// @router / [get]
func (req *ManufacturerController) Get() {
	var results []*things.Manufacturer
	man := things.Manufacturer{}
	err := things.Ctx.Manufacturer.Read(&man, &results)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": results}
	}

	req.ServeJSON()
}

// @Title SaveManufacturer
// @Description Saves a new manufacturer
// @Param	body		body 	things.Manufacturer	true		"body for manufacturer"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *ManufacturerController) Post() {
	var obj things.Manufacturer
	json.Unmarshal(req.Ctx.Input.RequestBody, &obj)

	_, err := things.Ctx.Manufacturer.Create(&obj)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "Save Successful."}
	}

	req.ServeJSON()
}
