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

func (ctrl *SecureController) ServeBinary(data []byte, filename string) {
	output := ctrl.Ctx.Output

	output.Header("Content-Description", "File Transfer")
	output.Header("Content-Type", "application/octet-stream")
	output.Header("Content-Disposition", "attachment; filename="+filename)
	output.Header("Content-Transfer-Encoding", "binary")
	output.Header("Expires", "0")
	output.Header("Cache-Control", "must-revalidate")
	output.Header("Pragma", "public")

	output.Body(data)
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
