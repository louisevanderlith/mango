package controllers

import (
	"github.com/louisevanderlith/mango/pkg/control"
)

type CreateController struct {
	control.UIController
}

func NewCreateCtrl(ctrlMap *control.ControllerMap) *CreateController {
	result := &CreateController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *CreateController) Get() {
	c.Setup("create")
}

func (c *CreateController) GetStep() {
	step := c.Ctx.Input.Param(":step")
	c.Setup(step)

	c.Data["StepNo"] = step
}

// POST must start ad upload and verification
func (c *CreateController) Post() {

	// Verify VIN
	// Upload Photos
	// Confirm Vehicle Match
	// Do something with tag...
	// Save Object
}
