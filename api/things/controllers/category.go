package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db/things"
	"github.com/louisevanderlith/mango/util/enums"
)

type CategoryController struct {
	util.SecureController
}

func init() {
	auths := make(util.ActionAuth)
	auths["GET"] = enums.User

	util.ProtectMethods(auths)
}

func (req *CategoryController) Get() {
	var results []*things.Category
	cat := things.Category{}
	err := things.Ctx.Category.Read(cat, &results)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": results}
	}

	req.ServeJSON()
}
