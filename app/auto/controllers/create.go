package controllers

import (
	"fmt"

	"github.com/louisevanderlith/mango/util/control"
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
	fmt.Printf("GETTING A STEP! On step %s\n", step)

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
