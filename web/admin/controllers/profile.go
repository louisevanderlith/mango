package controllers

import (
	"log"

	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/mango/pkg/control"

	"github.com/louisevanderlith/mango/app/admin/logic"
)

type ProfileController struct {
	control.UIController
}

func NewProfileCtrl(ctrlMap *control.ControllerMap) *ProfileController {
	result := &ProfileController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *ProfileController) Get() {
	c.Setup("profile")
	c.CreateSideMenu(logic.GetMenu("/profile"))

	data, err := logic.GetSites(c.GetInstanceID())

	c.Serve(data, err)
}

func (c *ProfileController) GetEdit() {
	c.Setup("profileEdit")
	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		log.Print("GetEdit:", err)
	}

	data, err := logic.GetSite(key, c.GetInstanceID())

	c.Serve(data, err)
}
