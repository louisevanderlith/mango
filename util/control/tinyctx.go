package control

import (
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

type tinyCtx struct {
	ApplicationName string
	RequiredRole    enums.RoleType
	URL             string
	Method          string
	SessionID       string
	Service         *util.Service
}

const avosession = "avosession"

func newTinyCtx(m *ControllerMap, ctx *context.Context) *tinyCtx {
	result := tinyCtx{}

	url, token := removeToken(ctx.Request.RequestURI)

	if token == "" {
		token = ctx.GetCookie(avosession)
	}

	actMethod := strings.ToUpper(ctx.Request.Method)
	required := m.GetRequiredRole(url, actMethod)

	result.RequiredRole = required
	result.ApplicationName = m.service.Name
	result.URL = url
	result.Method = actMethod
	result.SessionID = token
	result.Service = m.service

	return &result
}

func (ctx *tinyCtx) allowed() bool {
	return ctx.hasRole(ctx.RequiredRole)
}

func (ctx *tinyCtx) getUserKey() husk.Key {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return husk.CrazyKey()
	}

	return cookie.UserKey
}

func (ctx *tinyCtx) getUsername() string {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return "Unknown"
	}

	return cookie.Username
}

func (ctx *tinyCtx) getIP() string {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return "-1.-1.-1.-1"
	}

	return cookie.IP
}

func (ctx *tinyCtx) getLocation() string {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return "Uknown"
	}

	return cookie.Location
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

func (ctx *tinyCtx) hasRole(required enums.RoleType) bool {
	role := ctx.getRole()

	return role <= required
}

//TODO: use channels
//getAvoCookie also checks cookie validity, so repeated calls are required
func (ctx *tinyCtx) getAvoCookie() (*Cookies, error) {
	resp := util.GETMessage(ctx.Service.ID, "Secure.API", "login", "avo", ctx.SessionID)

	if resp.Failed() {
		return nil, resp
	}

	result := resp.Data.(Cookies)

	return &result, nil
}

func removeToken(url string) (cleanURL, token string) {
	idx := strings.LastIndex(url, "?token")
	tokenIdx := strings.LastIndex(url, "=") + 1

	if idx == -1 {
		return "/", ""
	}

	cleanURL = url[:idx]
	token = url[tokenIdx:]

	return cleanURL, token
}
