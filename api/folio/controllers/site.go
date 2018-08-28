package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/louisevanderlith/mango/core/folio"
	"github.com/louisevanderlith/mango/util/control"
)

type SiteController struct {
	control.APIController
}

// @Title RegisterWebsite
// @Description Register a Website
// @Param	body		body 	folio.Profile	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *SiteController) Post() {
	var site folio.Profile
	json.Unmarshal(req.Ctx.Input.RequestBody, &site)

	id, err := folio.Ctx.Profiles.Create(&site)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": id}
	}

	req.ServeJSON()
}

// @Title UpdateWebsite
// @Description Updates a Website
// @Param	body		body 	folio.Profile	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *SiteController) Put() {
	var site folio.Profile
	json.Unmarshal(req.Ctx.Input.RequestBody, &site)
	err := folio.Ctx.Profiles.Update(&site)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "Website has been updated."}
	}

	req.ServeJSON()
}

// @Title GetSites
// @Description Gets all sites
// @Success 200 {[]folio.Profile} []folio.Portfolio]
// @router / [get]
func (req *SiteController) Get() {
	if req.Ctx.Output.Status != 401 {
		var results folio.Profiles
		prof := folio.Profile{}
		err := folio.Ctx.Profiles.Read(&prof, &results)

		if err != nil {
			req.Ctx.Output.SetStatus(500)
			req.Data["json"] = map[string]string{"Error": err.Error()}
		} else {
			req.Data["json"] = map[string]interface{}{"Data": results}
		}
	}

	req.ServeJSON()
}

// @Title GetSite
// @Description Gets customer website/profile
// @Param	site			path	string 	true		"customer website name OR ID"
// @Success 200 {folio.Profile} folio.Profile
// @router /:site [get]
func (req *SiteController) GetOne() {
	siteParam := req.Ctx.Input.Param(":site")
	msg := folio.Profile{}

	if id, err := strconv.ParseInt(siteParam, 10, 32); err == nil {
		msg.Id = id
	} else {
		msg.Title = siteParam
	}

	result, err := folio.GetPortfolio()
	//folio.Ctx.Profiles.ReadOne(&msg, "SocialLinks", "PortfolioItems", "AboutSections", "Headers")

	req.Serve(err, result)
}
