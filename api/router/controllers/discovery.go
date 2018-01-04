package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/api/router/logic"
	"github.com/louisevanderlith/mango/util"
	"strconv"
	"github.com/louisevanderlith/mango/util/control"
)

type DiscoveryController struct {
	control.APIController
}

// @Title RegisterAPI
// @Description Register an API
// @Param	body		body 	util.Service	true		"body for service content"
// @Success 200 {string} models.Service.ID
// @Failure 403 body is empty
// @router / [post]
func (req *DiscoveryController) Post() {
	var service util.Service
	json.Unmarshal(req.Ctx.Input.RequestBody, &service)

	appID := logic.AddService(service)

	req.Data["json"] = map[string]string{"AppID": appID}
	req.ServeJSON()
}

// @Title GetService
// @Description Gets the recommended service
// @Param	appID			path	string 	true		"the application requesting a service"
// @Param	serviceName		path 	string	true		"the name of the service you want to get"
// @Param	clean			path 	bool	false		"clean will return a user friendly URL and not the application's actual URL"
// @Success 200 {string} util.Service.URL
// @Failure 403 :serviceName or :appID is empty
// @router /:appID/:serviceName/:clean [get]
func (req *DiscoveryController) Get() {
	appID := req.Ctx.Input.Param(":appID")
	serviceName := req.Ctx.Input.Param(":serviceName")
	clean, cleanErr := strconv.ParseBool(req.Ctx.Input.Param(":clean"))

	if cleanErr != nil {
		clean = false
	}

	if appID != "" && serviceName != "" {
		url, err := logic.GetServicePath(serviceName, appID, clean)

		if err != nil {
			req.Ctx.Output.SetStatus(500)
			req.Data["json"] = err.Error()
		} else {
			req.Data["json"] = url
		}
	}

	req.ServeJSON()
}

// @Title GetDirtyService
// @Description Gets the recommended service
// @Param	appID			path	string 	true		"the application requesting a service"
// @Param	serviceName		path 	string	true		"the name of the service you want to get"
// @Success 200 {string} util.Service.URL
// @Failure 403 :serviceName or :appID is empty
// @router /:appID/:serviceName [get]
func (req *DiscoveryController) GetDirty() {
	appID := req.Ctx.Input.Param(":appID")
	serviceName := req.Ctx.Input.Param(":serviceName")

	if appID != "" && serviceName != "" {
		url, err := logic.GetServicePath(serviceName, appID, false)

		if err != nil {
			req.Ctx.Output.SetStatus(500)
			req.Data["json"] = err.Error()
		} else {
			req.Data["json"] = url
		}
	}

	req.ServeJSON()
}
