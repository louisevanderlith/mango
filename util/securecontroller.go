package util

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/util/enums"
	"net/http"
)

type SecureController struct {
	beego.Controller
}

type ActionAuth map[string]enums.RoleType

var authFunctions ActionAuth

func init() {
	authFunctions = make(ActionAuth)
}

func (ctrl *SecureController) Prepare() {
	if !userAllowed(ctrl) {
		ctrl.Ctx.Output.SetStatus(http.StatusUnauthorized)
		ctrl.Data["json"] = map[string]string{"Error": "User not authorized to access this content."}
	}
}

func (ctrl *SecureController) GetAvoToken() string {
	return ctrl.Ctx.GetCookie("avotoken")
}

func (ctrl *SecureController) ExpireAvoToken() {
	ctrl.Ctx.SetCookie("avotoken", "expired", 0)
}

func (ctrl *SecureController) SetAvoToken(token string) {
	ctrl.Ctx.SetCookie("avotoken", token, 600, "/", "avosa.co.za", true, true)
}

func ProtectMethods(auths ActionAuth) {
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