package control

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

type tinyCtx struct {
	ApplicationName string
	RequiredRole    enums.RoleType
	Controller      string
	Method          string
	SessionID       string
}

const avosession = "avosession"

func newTinyCtx(ctx *context.Context) *tinyCtx {
	result := tinyCtx{}

	ctrl, sess := removeToken(ctx.Request.RequestURI)

	if sess == "" {
		sess = ctx.GetCookie(avosession)
	}

	actMethod := strings.ToUpper(ctx.Request.Method)

	result.RequiredRole = GetRequiredRole(ctrl, actMethod)
	result.ApplicationName = controllerMap.applicationName
	result.Controller = ctrl
	result.Method = actMethod
	result.SessionID = sess

	return &result
}

func (ctx *tinyCtx) allowed() bool {
	return ctx.hasRole(ctx.RequiredRole)
}

func (ctx *tinyCtx) getUserID() int64 {
	result := int64(-1)

	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return result
	}

	result = cookie.UserID

	return result
}

func (ctx *tinyCtx) getUsername() string {
	result := "Unknown"

	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return result
	}

	result = cookie.Username

	return result
}

func (ctx *tinyCtx) getRole() enums.RoleType {
	result := enums.Unknown

	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return result
	}

	appName := ctx.ApplicationName
	if role, ok := cookie.UserRoles[appName]; ok {
		result = role
	}

	return result
}

func (ctx *tinyCtx) hasRole(actionRole enums.RoleType) bool {
	role := ctx.getRole()

	return role < actionRole
}

//TODO: use channels
//getAvoCookie also checks cookie validity, so repeated calls are required
func (ctx *tinyCtx) getAvoCookie() (result Cookies, finalError error) {
	contents, statusCode := util.GETMessage("Secure.API", "login", "avo", ctx.SessionID)
	data := util.MarshalToMap(contents)

	if statusCode != 200 {
		var dataErr string
		err := json.Unmarshal(*data["Error"], &dataErr)

		if err != nil {
			log.Println("getAvoCookie:", err)
		}

		finalError = errors.New(dataErr)
	} else {
		err := json.Unmarshal(*data["Data"], &result)

		if err != nil {
			log.Println("getAvoCookie:", err)
		}
	}

	return result, finalError
}

func removeToken(url string) (cleanURL, token string) {
	idx := strings.LastIndex(url, "?token")
	tokenIdx := strings.LastIndex(url, "=") + 1

	if idx != -1 {
		token = url[tokenIdx:]
		cleanURL = url[:idx]
	}

	if cleanURL == "" {
		cleanURL = "/"
	}

	return cleanURL, token
}
