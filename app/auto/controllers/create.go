package controllers

import (
	"github.com/louisevanderlith/mango/util/control"
)

type CreateController struct {
	control.UIController
}

func (c *CreateController) Get() {
	c.Setup("create")
}

// POST must start ad upload and verification
func (c *CreateController) Post() {

	// Verify VIN
	// Upload Photos
	// Confirm Vehicle Match
	// Do something with tag...
	// Save Object
}
