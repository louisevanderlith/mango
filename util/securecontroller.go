package util

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/util/enums"
)

type SecureController struct{
	beego.Controller
}

var authFunctions map[string]enums.RoleType

func init() {
	authFunctions = make(map[string]enums.RoleType)
}

func (ctrl *SecureController) Prepare() {

	if !userAllowed(ctrl) {
		ctrl.CustomAbort(401, "You don't have permission to access this content.")
	}
}

func ProtectMethods(auths map[string]enums.RoleType) {
	authFunctions = auths
}

func userAllowed(ctrl *SecureController) bool {
	result := true
	method := ctrl.Ctx.Request.Method
	authFunc, hasKey := authFunctions[method]

	if hasKey {
		userSession := ctrl.Ctx.Request.Header.Get("avotoken")

		result = hasRole(userSession, authFunc)
	}

	return result
}