package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/api/router/logic"
)

type DiscoveryController struct {
	beego.Controller
}

// @Title RegisterAPI
// @Description Register an API
// @Param	body		body 	models.Service	true		"body for service content"
// @Success 200 {string} models.Service.ID
// @Failure 403 body is empty
// @router / [post]
func (req *DiscoveryController) Post() {
	var service logic.Service
	json.Unmarshal(req.Ctx.Input.RequestBody, &service)

	appID := logic.AddService(service)

	req.Data["json"] = map[string]string{"AppID": appID}
	req.ServeJSON()
}

// @Title GetService
// @Description Gets the recommended service
// @Param	appID			path	string 	true		"the application requesting a service"
// @Param	serviceName		path 	string	true		"the name of the service you want to get"
// @Success 200 {string} models.Service.URL
// @Failure 403 :serviceName or :appID is empty
// @router /:appID/:serviceName [get]
func (req *DiscoveryController) Get() {
	appID := req.Ctx.Input.Param(":appID")
	serviceName := req.Ctx.Input.Param(":serviceName")

	if appID != "" && serviceName != "" {
		url, err := logic.GetServicePath(serviceName, appID)

		if err != nil {
			req.Data["json"] = err.Error()
		} else {
			req.Data["json"] = url
		}
	}

	req.ServeJSON()
}
