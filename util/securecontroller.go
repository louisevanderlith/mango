package util

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/util/enums"
	"net/http"
	"encoding/json"
	"log"
	"errors"
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
		result = hasRole(authFunc)
	}

	return result
}

func hasRole(funcRole enums.RoleType) bool {
	result := false
	roles, err := loadRoles()

	if err == nil {
		for _, val := range roles {
			if val <= funcRole {
				result = true
				break
			}
		}
	}

	return result
}

func loadRoles() ([]enums.RoleType, error) {
	var result []enums.RoleType
	var finalError error

	contents, statusCode := GETMessage("Security.API", "session")
	data := MarshalToMap(contents)

	if statusCode != 200 {
		var dataErr string
		err := json.Unmarshal(*data["Error"], &dataErr)

		if err != nil {
			log.Printf("loadRoles: ", err)
		}

		finalError = errors.New(dataErr)
	} else {
		err := json.Unmarshal(*data["Data"], &result)

		if err != nil {
			log.Printf("loadRoles: ", err)
		}
	}

	return result, finalError
}
