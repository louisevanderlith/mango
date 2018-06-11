package controllers

import (
	"github.com/louisevanderlith/mango/app/admin/logic"
)

type CommsController struct {
	logic.MenuController
}

func (c *CommsController) Get() {
	c.Setup("comms")

	data, err := logic.GetCommsMessages()

	c.Serve(err, data)
}
