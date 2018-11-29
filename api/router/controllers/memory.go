package controllers

import (
	"github.com/louisevanderlith/mango/api/router/logic"
	"github.com/louisevanderlith/mango/util/control"
)

type MemoryController struct {
	control.APIController
}

func NewMemoryCtrl(ctrlMap *control.ControllerMap) *MemoryController {
	result := &MemoryController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetStuff
// @Description Gets the stuff registered
// @Success 200 {string} models.Service.ID
// @router / [get]
func (req *MemoryController) Get() {
	srvMap := logic.GetServiceMap()
	req.Serve(srvMap, nil)
}
