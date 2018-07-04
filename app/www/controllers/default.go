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

func getTopMenu() []control.Menu {
	var result []control.Menu

	result = append(result, control.Menu{
		Name:  "Portfolio",
		Class: "home gome fa-home",
		Link:  "#portfolio",
	})

	result = append(result, control.Menu{
		Name:  "About Us",
		Class: "globe world",
		Link:  "#aboutus",
	})

	result = append(result, control.Menu{
		Name:  "Contact",
		Class: "globe world",
		Link:  "#contact",
	})

	/*result = append(result, control.Menu{
		Name:  "Has Children",
		Class: "pregnant",
		Children: []control.Menu{
			control.Menu{
				Name:  "Home",
				Class: "home gome fa-home",
			},
		},
	})*/

	return result
}
