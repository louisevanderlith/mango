package util

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/util/enums"
)

type SecureController struct {
	beego.Controller
}

var authFunctions map[string]enums.RoleType

func init() {
	authFunctions = make(map[string]enums.RoleType)
}

func (ctrl *SecureController) Prepare() {
	if !userAllowed(ctrl) {
		ctrl.Ctx.Output.SetStatus(401)
		ctrl.Data["json"] = map[string]string{"Error": "User not authorized to access this content."}
	}
}

func (ctrl *SecureController) GetAvoToken() string {
	return ctrl.Ctx.GetCookie("avotoken");
}

func (ctrl *SecureController) ExpireAvoToken() {
	ctrl.Ctx.SetCookie("avotoken", "expired", 0)
}

func (ctrl *SecureController) SetAvoToken(token string) {
	ctrl.Ctx.SetCookie("avotoken", token, 600, "/", "avosa.co.za", true, true)
}

func ProtectMethods(auths map[string]enums.RoleType) {
	authFunctions = auths
}

func userAllowed(ctrl *SecureController) bool {
	result := true
	method := ctrl.Ctx.Request.Method
	authFunc, hasKey := authFunctions[method]

	if hasKey {
		userSession := ctrl.GetAvoToken()

		result = hasRole(userSession, authFunc)
	}

	return result
}