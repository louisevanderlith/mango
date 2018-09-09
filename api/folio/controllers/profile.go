package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/mango/core/folio"
	"github.com/louisevanderlith/mango/util/control"
)

type ProfileController struct {
	control.APIController
}

// @Title RegisterWebsite
// @Description Register a Website
// @Param	body		body 	folio.Profile	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *ProfileController) Post() {
	var site folio.Profile
	json.Unmarshal(req.Ctx.Input.RequestBody, &site)

	rec, err := site.Create()

	req.Serve(err, rec)
}

// @Title UpdateWebsite
// @Description Updates a Website
// @Param	body		body 	folio.Profile	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *ProfileController) Put() {
	with, err := req.GetKeyedRequest()

	if err != nil {
		req.Serve(err, nil)
		return
	}

	body := with.Body.(folio.Profile)
	err = body.Update(with.Key)

	req.Serve(err, nil)
}

// @Title GetSites
// @Description Gets all sites
// @Success 200 {[]folio.Profile} []folio.Portfolio]
// @router /:pageData(^[A-Z](?:_?[0-9]+)*$) [get]
func (req *ProfileController) Get() {
	page, size := req.GetPageData()

	results, err := folio.GetProfiles(page, size)

	req.Serve(err, results)
}

// @Title GetSite
// @Description Gets customer website/profile
// @Param	site			path	string 	true		"customer website name OR ID"
// @Success 200 {folio.Profile} folio.Profile
// @router /:site [get]
func (req *ProfileController) GetOne() {
	siteParam := req.Ctx.Input.Param(":site")

	var result *folio.Profile
	var err error

	if key := husk.ParseKey(siteParam); key != husk.CrazyKey() {
		result, err = folio.GetProfile(key)
	} else {
		result, err = folio.GetProfileByName(siteParam)
	}

	req.Serve(err, result)
}
