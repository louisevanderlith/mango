package controllers

import (
	"github.com/louisevanderlith/mango/pkg/control"
)

type CreditController struct {
	control.APIController
}

func NewCreditCtrl(ctrlMap *control.ControllerMap) *CreditController {
	result := &CreditController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetCreditBalance
// @Description Gets the current user's credit balance.
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [get]
func (req *CreditController) Get() {

}
