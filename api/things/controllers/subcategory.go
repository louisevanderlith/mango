package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db/things"
	"github.com/louisevanderlith/mango/util/enums"
	"encoding/json"
)

type SubCategoryController struct{
	util.SecureController
}

func init(){
	auths := make(util.ActionAuth)
	auths["GET"] = enums.User

	util.ProtectMethods(auths)
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

func (req *SubCategoryController) Post() {
	var obj things.SubCategory
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
