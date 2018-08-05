package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/core/folio"
	"github.com/louisevanderlith/mango/util/control"
)

type AboutController struct {
	control.APIController
}

// @Title CreateAboutSection
// @Description Creates an about section for a current site.
// @Param	body		body 	folio.About	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *AboutController) Post() {
	var about folio.About
	json.Unmarshal(req.Ctx.Input.RequestBody, &about)

	_, err := folio.Ctx.Abouts.Create(&about)

	req.Serve(err, "About Section has been created.")
}

// @Title UpdateAbout
// @Description Updates a About section on a current site
// @Param	body		body 	folio.About	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *AboutController) Put() {
	var about folio.About
	json.Unmarshal(req.Ctx.Input.RequestBody, &about)

	err := folio.Ctx.Abouts.Update(&about)

	req.Serve(err, "About Section has been updated.")
}
