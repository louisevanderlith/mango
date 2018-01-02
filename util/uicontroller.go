package util

import (
	"net/http"
	"fmt"
	"net/url"
)

type UIController struct {
	SecureController
	HasScript  bool
	ScriptName string
}

func (ctrl *UIController) Prepare() {
	ctrl.SecureController.Prepare()

	if ctrl.Ctx.Output.Status == 401 {
		securityURL, err := GetServiceURL("Security.API", true)

		if err == nil {
			req := ctrl.Ctx.Request
			moveURL := fmt.Sprintf("%s://%s%s", ctrl.Ctx.Input.Scheme(), req.Host, req.RequestURI)
			loginURL := fmt.Sprintf("%sv1/login?return=%s", securityURL, url.QueryEscape(moveURL))

			ctrl.Redirect(loginURL, http.StatusTemporaryRedirect)
		}
	}

	ctrl.Layout = "master.html"
}

func (ctrl *UIController) Setup(name string) {
	ctrl.TplName = "content/" + name + ".html"

	// By default we want to include scripts
	// Set this to false in your controller, when scripts aren't needed
	ctrl.Data["HasScript"] = true
	ctrl.Data["ScriptName"] = name + ".entry.js"
	ctrl.Data["InstanceKey"] = GetInstanceKey()
}
