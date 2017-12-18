package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
	"github.com/louisevanderlith/mango/db/folio"
	"encoding/json"
)

type SiteController struct {
	util.SecureController
}

func init() {
	auths := make(util.ActionAuth)
	auths["POST"] = enums.Admin

	util.ProtectMethods(auths)
}

// @Title Register Website
// @Description Register a Website
// @Param	body		body 	models.Service	true		"body for service content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *SiteController) Post() {
	var site folio.Profile
	json.Unmarshal(req.Ctx.Input.RequestBody, &site)

	_, err := folio.Ctx.Profile.Create(&site)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "Website has been created."}
	}

	req.ServeJSON()
}

// @Title GetSite
// @Description Gets customer website/profile
// @Param	siteName			path	string 	true		"customer website name"
// @Success 200 {string} string
// @router /:siteName [get]
func (req *SiteController) Get() {
	if req.Ctx.Output.Status != 401 {
		siteName := req.Ctx.Input.Param(":siteName")
		msg := folio.Profile{}
		msg.Title = siteName

		result, err := folio.Ctx.Profile.ReadOne(&msg, "SocialLinks", "PortfolioItems", "AboutSections")

		if err != nil {
			req.Ctx.Output.SetStatus(500)
			req.Data["json"] = map[string]string{"Error": err.Error()}
		} else {
			req.Data["json"] = map[string]interface{}{"Data": result}
		}
	}

	req.ServeJSON()
}
