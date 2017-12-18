package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db/things"
	"github.com/louisevanderlith/mango/util/enums"
	"encoding/json"
)

type ManufacturerController struct {
	util.SecureController
}

func init() {
	auths := make(util.ActionAuth)
	auths["GET"] = enums.User

	util.ProtectMethods(auths)
}

func (req *ManufacturerController) Get() {
	var results []*things.Manufacturer
	man := things.Manufacturer{}
	err := things.Ctx.Manufacturer.Read(man, &results)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": results}
	}

	req.ServeJSON()
}

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
