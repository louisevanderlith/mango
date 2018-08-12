package controllers

import (
	"fmt"

	"github.com/louisevanderlith/mango/util/control"
)

type CreateController struct {
	control.UIController
}

func (c *CreateController) Get() {
	c.Setup("create")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
}

func (c *CreateController) GetStep() {
	step := c.Ctx.Input.Param(":step")
	c.Setup(step)
	fmt.Printf("GETTING A STEP! On step %s\n", step)

	c.Data["StepNo"] = step
}
