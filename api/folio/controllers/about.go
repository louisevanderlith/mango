package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/db/folio"
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

	_, err := folio.Ctx.About.Create(&about)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "About Section has been created."}
	}

	req.ServeJSON()
}
