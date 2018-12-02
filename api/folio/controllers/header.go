package controllers

import (
	"github.com/louisevanderlith/mango/core/folio"
	"github.com/louisevanderlith/mango/pkg/control"
)

type HeaderController struct {
	control.APIController
}

func NewHeaderCtrl(ctrlMap *control.ControllerMap) *HeaderController {
	result := &HeaderController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title CreateHeaderItem
// @Description Creates a Portfolio Item on a current site
// @Param	body		body 	folio.Portfolio	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *HeaderController) Post() {
	with, err := req.GetKeyedRequest()

	if err != nil {
		req.Serve(nil, err)
		return
	}

	err = folio.AddHeaderSection(with.Key, with.Body.(folio.Header))

	req.Serve(nil, err)
}
