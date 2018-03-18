package controllers

import (
	"github.com/louisevanderlith/mango/util/control"
)

type DefaultController struct {
	control.UIController
}

func (c *DefaultController) Get() {
	c.Setup("default")
	/*siteName := c.Ctx.Input.Param(":siteName")
	data, err := logic.GetProfileSite(siteName)

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["error"] = err
	} else {
		c.Data["data"] = data
	}*/
}
