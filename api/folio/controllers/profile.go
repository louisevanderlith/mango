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

func NewProfileCtrl(ctrlMap *control.ControllerMap) *ProfileController {
	result := &ProfileController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetSites
// @Description Gets all sites
// @Success 200 {[]folio.Profile} []folio.Portfolio]
// @router /:pageData^[A-Z]+:[0-9]+$ [get]
func (req *ProfileController) Get() {
	page, size := req.GetPageData()

	results := folio.GetProfiles(page, size)

	req.Serve(results, nil)
}

// @Title GetSite
// @Description Gets customer website/profile
// @Param	site			path	string 	true		"customer website name OR ID"
// @Success 200 {folio.Profile} folio.Profile
// @router /:site [get]
func (req *ProfileController) GetOne() {
	siteParam := req.Ctx.Input.Param(":site")

	key, err := husk.ParseKey(siteParam)

	if err != nil && key == nil {
		byName, err := folio.GetProfileByName(siteParam)
		req.Serve(byName, err)
		return
	}

	result, err := folio.GetProfile(key)

	req.Serve(result, err)
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

	rec := site.Create()

	req.Serve(rec, nil)
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
		req.Serve(nil, err)
		return
	}

	body := with.Body.(folio.Profile)
	err = body.Update(with.Key)

	req.Serve(nil, err)
}
