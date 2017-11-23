package controllers

import (
	"github.com/louisevanderlith/mango/util"
)

type MainController struct {
	util.UIController
}

func (c *MainController) Get() {
	c.Setup("main")

}
