package util

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/util/enums"
)

type BaseController struct {
	beego.Controller
	HasScript  bool
	ScriptName string
}

var authFunctions map[string]enums.RoleType

func init() {
	authFunctions = make(map[string]enums.RoleType)
}

// Prepare is a virtual function called by beego before each Controller function
func (this *BaseController) Prepare() {
	this.Layout = "master.html"

	if !userAllowed(this) {
		this.CustomAbort(401, "You don't have permission to access this content.")
	}
}

func ProtectMethods(auths map[string]enums.RoleType) {
	authFunctions = auths
}

func (ctrl *BaseController) Setup(name string) {
	ctrl.TplName = "content/" + name + ".html"

	// By default we want to include scripts
	// Set this to false in your controller, when scripts aren't needed
	ctrl.Data["HasScript"] = true
	ctrl.Data["ScriptName"] = name + ".entry.js"
}

func userAllowed(ctrl *BaseController) bool {
	result := true
	method := ctrl.Ctx.Request.Method
	authFunc, hasKey := authFunctions[method]

	if hasKey {
		userSession := ctrl.Ctx.Request.Header.Get("avotoken")
		roles := getUserRoles(userSession)

		result = hasRole(roles, authFunc)
	}

	return result
}

func getUserRoles(token string) []enums.RoleType {
	var result []enums.RoleType

	return result
}

func hasRole(roles []enums.RoleType, funcRole enums.RoleType) bool {
	result := false

	for _, val := range roles {
		if val == funcRole {
			result = true
			break
		}
	}

	return result
}
