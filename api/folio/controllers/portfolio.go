package controllers

import (
	"github.com/louisevanderlith/mango/core/folio"
	"github.com/louisevanderlith/mango/util/control"
)

type PortfolioController struct {
	control.APIController
}

func NewPortfolioCtrl(ctrlMap *control.ControllerMap) *PortfolioController {
	result := &PortfolioController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title CreatePortfolioItem
// @Description Creates a Portfolio Item on a current site
// @Param	body		body 	folio.Portfolio	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *PortfolioController) Post() {
	with, err := req.GetKeyedRequest()

	if err != nil {
		req.Serve(err, nil)
		return
	}

	err = folio.AddPortfolioSection(with.Key, with.Body.(folio.Portfolio))

	req.Serve(err, nil)
}
