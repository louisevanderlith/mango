package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
	"github.com/louisevanderlith/mango/app/admin/logic"
)

type ModelController struct {
	util.UIController
}

func init() {
	auths := make(util.ActionAuth)
	auths["GET"] = enums.Admin
	auths["POST"] = enums.Admin

	util.ProtectMethods(auths)
}

func (c *ModelController) Get() {
	c.Setup("model")

	data, err := logic.GetModels()

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["error"] = err
	} else {
		c.Data["data"] = data
	}
}

func (c *ModelController) Post() {

}
