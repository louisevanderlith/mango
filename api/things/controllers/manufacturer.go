package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db/things"
)

type ManufacturerController struct{
	util.SecureController
}

func (req *ManufacturerController) Get(){
	man := things.Manufacturer{}
	results, err := man.ReadAll()

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": results}
	}

	req.ServeJSON()
}