package controllers

import (
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
	with, err := req.GetKeyedRequest()

	if err != nil {
		req.Serve(err, nil)
		return
	}

	err = folio.AddAboutSection(with.Key, with.Body.(folio.About))

	req.Serve(err, nil)
}
