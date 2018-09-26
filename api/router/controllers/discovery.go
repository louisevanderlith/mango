package controllers

import (
	"encoding/json"
	"errors"
	"log"

	"strconv"

	"github.com/louisevanderlith/mango/api/router/logic"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/control"
)

type DiscoveryController struct {
	control.APIController
}

func NewDiscoveryCtrl(ctrlMap *control.ControllerMap) *DiscoveryController {
	result := &DiscoveryController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title RegisterAPI
// @Description Register an API
// @Param	body		body 	util.Service	true		"body for service content"
// @Success 200 {string} models.Service.ID
// @Failure 403 body is empty
// @router / [post]
func (req *DiscoveryController) Post() {
	service := &util.Service{}
	json.Unmarshal(req.Ctx.Input.RequestBody, service)

	log.Printf("SERV:%+v\n", service)
	appID, err := logic.AddService(service)

	req.Serve(appID, err)
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

	if appID == "" || serviceName == "" {
		err := errors.New("appID AND serviceName must be populated")
		req.Serve(nil, err)
		return
	}

	clean, cleanErr := strconv.ParseBool(req.Ctx.Input.Param(":clean"))

	if cleanErr != nil {
		clean = false
	}

	url, err := logic.GetServicePath(serviceName, appID, clean)
	req.Serve(url, err)
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

	if appID == "" || serviceName == "" {
		err := errors.New("appID AND serviceName must be populated")
		req.Serve(nil, err)
		return
	}

	url, err := logic.GetServicePath(serviceName, appID, false)

	req.Serve(url, err)
}
