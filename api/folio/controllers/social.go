package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/db/folio"
	"github.com/louisevanderlith/mango/util/control"
)

type SocialController struct {
	control.APIController
}

// @Title CreateSocialLink
// @Description Creates a Social Link on a current site
// @Param	body		body 	folio.SocialLink	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *SocialController) Post() {
	var link folio.SocialLink
	json.Unmarshal(req.Ctx.Input.RequestBody, &link)

	_, err := folio.Ctx.SocialLink.Create(&link)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "Social Media Item has been created."}
	}

	req.ServeJSON()
}
