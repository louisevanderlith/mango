package controllers

import (
	"github.com/louisevanderlith/mango/app/www/logic"
	"github.com/louisevanderlith/mango/pkg/control"
)

type DefaultController struct {
	control.UIController
}

func NewDefaultCtrl(ctrlMap *control.ControllerMap) *DefaultController {
	result := &DefaultController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *DefaultController) Get() {
	c.Setup("default")
	c.CreateTopMenu(getTopMenu())
	siteName := c.Ctx.Input.Param(":siteName")
	data, err := logic.GetProfileSite(c.GetInstanceID(), siteName)

	c.Serve(data, err)
}

func getTopMenu() *control.Menu {
	result := control.NewMenu("/home")

	result.AddItem("#portfolio", "Portfolio", "home gome fa-home", nil)
	result.AddItem("#aboutus", "About Us", "home gome fa-home", nil)
	result.AddItem("#contact", "Contact", "home gome fa-home", nil)

	return result
}
