package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
	"github.com/louisevanderlith/mango/app/admin/logic"
)

type SiteController struct {
	util.UIController
}

func init(){
	auths := make(util.ActionAuth)
	auths["GET"] = enums.Admin
	auths["POST"] = enums.Admin

	util.ProtectMethods(auths)
}

func (c *SiteController) Get() {
	c.Setup("site")

	data, err := logic.GetSites()

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["error"] = err
	} else {
		c.Data["data"] = data
	}
}
