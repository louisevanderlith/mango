package controllers

import (
	"github.com/louisevanderlith/mango/core/folio"
	"github.com/louisevanderlith/mango/util/control"
)

type SocialController struct {
	control.APIController
}

func NewSocialCtrl(ctrlMap *control.ControllerMap) *SocialController {
	result := &SocialController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title CreateSocialLink
// @Description Creates a Social Link on a current site
// @Param	body		body 	folio.SocialLink	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *SocialController) Post() {
	with, err := req.GetKeyedRequest()

	if err != nil {
		req.Serve(err, nil)
		return
	}

	err = folio.AddSocialLink(with.Key, with.Body.(folio.SocialLink))

	req.Serve(err, nil)
}
