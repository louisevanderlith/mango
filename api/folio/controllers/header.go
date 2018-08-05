package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/core/folio"
	"github.com/louisevanderlith/mango/util/control"
)

type HeaderController struct {
	control.APIController
}

// @Title CreateHeaderItem
// @Description Creates a Portfolio Item on a current site
// @Param	body		body 	folio.Portfolio	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *HeaderController) Post() {
	var header folio.Header
	json.Unmarshal(req.Ctx.Input.RequestBody, &header)

	_, err := folio.Ctx.Headers.Create(&header)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "Header Item has been created."}
	}

	req.ServeJSON()
}

// @Title UpdateHeader
// @Description Updates a Header on a current site
// @Param	body		body 	folio.Header	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *HeaderController) Put() {
	var head folio.Header
	json.Unmarshal(req.Ctx.Input.RequestBody, &head)

	err := folio.Ctx.Headers.Update(&head)

	req.Serve(err, "Header has been updated.")
}
