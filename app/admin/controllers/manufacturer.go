package controllers

import "github.com/louisevanderlith/mango/util"

type ManufacturerController struct{
	util.UIController
}

func (c *ManufacturerController) Get(){
	c.Setup("manufacturer")
}
