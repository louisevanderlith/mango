package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/app/admin/logic"
)

type CommsController struct {
	util.BaseController
}

func (c *CommsController) Get() {
	c.Setup("comms")
	data, err := logic.GetCommsMessages()

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["error"] = err
	} else {
		c.Data["data"] = data
	}
}
