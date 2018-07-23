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
	c.CreateTopMenu(getTopMenu())
	siteName := c.Ctx.Input.Param(":siteName")
	data, err := logic.GetProfileSite(siteName)

	c.Serve(err, data)
}

func getTopMenu() *control.Menu {
	result := control.NewMenu("/home")

	result.AddItem("#portfolio", "Portfolio", "home gome fa-home", nil)
	result.AddItem("#aboutus", "About Us", "home gome fa-home", nil)
	result.AddItem("#contact", "Contact", "home gome fa-home", nil)

	return result
}
