package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db/things"
)

type CategoryController struct {
	util.SecureController
}

func (req *CategoryController) Get() {
	cat := things.Category{}
	results, err := cat.ReadAll()

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": results}
	}

	req.ServeJSON()
}
