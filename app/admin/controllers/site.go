package controllers

import (
	"log"
	"strconv"

	"github.com/louisevanderlith/mango/util/control"

	"github.com/louisevanderlith/mango/app/admin/logic"
)

type SiteController struct {
	control.UIController
}

func NewSiteCtrl(ctrlMap *control.ControllerMap) *SiteController {
	result := &SiteController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *SiteController) Get() {
	c.Setup("site")
	c.CreateSideMenu(logic.GetMenu("/site"))

	data, err := logic.GetSites(c.GetInstanceID())

	c.Serve(err, data)
}

func (c *SiteController) GetEdit() {
	c.Setup("siteEdit")
	id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)

	if err != nil {
		log.Print("GetEdit:", err)
	}

	data, err := logic.GetSite(id, c.GetInstanceID())

	c.Serve(err, data)
}
