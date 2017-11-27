package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db/things"
)

type SubCategoryController struct{
	util.SecureController
}

func (req *SubCategoryController) Get(){
	var results []*things.SubCategory
	scat := things.SubCategory{}
	err := things.Ctx.SubCategory.Read(scat, &results)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": results}
	}

	req.ServeJSON()
}
