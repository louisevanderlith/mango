package controllers

import (
	"github.com/louisevanderlith/mango/app/admin/logic"
	"github.com/louisevanderlith/mango/util/control"
)

type CommsController struct {
	control.UIController
}

func (c *CommsController) Get() {
	c.Setup("comms")
	c.CreateSideMenu(logic.GetMenu("/comms"))

	data, err := logic.GetCommsMessages()

	c.Serve(err, data)
}
