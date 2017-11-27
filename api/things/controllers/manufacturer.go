package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db/things"
)

type ManufacturerController struct{
	util.SecureController
}

func (req *ManufacturerController) Get(){
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