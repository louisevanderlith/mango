package controllers

import (
	"log"
	"strconv"

	"github.com/louisevanderlith/mango/app/admin/logic"
	"github.com/louisevanderlith/mango/util/control"
)

type SiteController struct {
	control.UIController
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

func (c *SiteController) GetEdit() {
	c.Setup("siteEdit")
	id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)

	if err != nil {
		log.Print("GetEdit:", err)
	}

	data, err := logic.GetSite(id)

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["error"] = err
	} else {
		c.Data["data"] = data
	}
}
