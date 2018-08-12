package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/core/folio"
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

	_, err := folio.Ctx.SocialLinks.Create(&link)

	req.Serve(err, "Social Media Item has been created.")
}

// @Title UpdateSocialLink
// @Description Updates a Social Link on a current site
// @Param	body		body 	folio.SocialLink	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *SocialController) Put() {
	var prtfolio folio.Portfolio

	folio.UpdateP(id, social)

	req.Serve(err, "Social Link has been updated.")
}
