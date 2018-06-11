package controllers

import (
	"github.com/louisevanderlith/mango/app/www/logic"
	"github.com/louisevanderlith/mango/util/control"
)

type DefaultController struct {
	control.UIController
}

func (c *DefaultController) Get() {
	c.Setup("default")
	siteName := c.Ctx.Input.Param(":siteName")
	data, err := logic.GetProfileSite(siteName)

	c.Serve(err, data)
}
