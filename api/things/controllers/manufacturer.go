package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/core/things"
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
	var results things.Manufacturers
	man := things.Manufacturer{}
	err := things.Ctx.Manufacturers.Read(&man, &results)

	req.Serve(err, results)
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

	_, err := things.Ctx.Manufacturers.Create(&obj)

	req.Serve(err, "Save Successful")
}
