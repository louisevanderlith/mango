package util

import (
	"github.com/louisevanderlith/mango/util/enums"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	HasScript  bool
	ScriptName string
}

var authFunctions map[string]enums.RoleType

func init(){
	authFunctions = make(map[string]enums.RoleType)
}

// Prepare is a virtual function called by beego before each Controller function
func (this *BaseController)Prepare(){
	this.Layout = "master.html"

	userAllowed :=  this.Ctx.
}

func (this *BaseController) Setup(name string) {
	this.TplName = "content/" + name + ".html"

	// By default we want to include scripts
	// Set this to false in your controller, when scripts aren't needed
	this.Data["HasScript"] = true
	this.Data["ScriptName"] = name + ".entry.js"
}

func userAllowed(ctrl *BaseController) bool {
	method := ctrl.Ctx.Request.Method
	authFunc, hasKey := authFunctions[method]

	if hasKey {
		userSession := ctrl.Ctx.Request.Header.Get("avotoken")
		role := 

	}

}
