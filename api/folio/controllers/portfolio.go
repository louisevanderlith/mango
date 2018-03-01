package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/db/folio"
	"github.com/louisevanderlith/mango/util/control"
)

type PortfolioController struct {
	control.APIController
}

// @Title CreatePortfolioItem
// @Description Creates a Portfolio Item on a current site
// @Param	body		body 	folio.Portfolio	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *PortfolioController) Post() {
	var portfolio folio.Portfolio
	json.Unmarshal(req.Ctx.Input.RequestBody, &portfolio)

	_, err := folio.Ctx.Portfolio.Create(&portfolio)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "Portfolio Item has been created."}
	}

	req.ServeJSON()
}
