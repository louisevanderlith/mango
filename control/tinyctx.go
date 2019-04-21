package control

import (
	"errors"
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

type TinyCtx struct {
	ApplicationName string
	RequiredRole    roletype.Enum
	URL             string
	Method          string
	SessionID       string
	Service         *mango.Service
}

const avosession = "avosession"

func findURLToken(ctx *context.Context) (string, string) {
	url, token := removeToken(ctx.Request.RequestURI)

	if token == "" {
		token = ctx.GetCookie(avosession)
	}

	return url, token
}

func NewTinyCtx(m *ControllerMap, ctx *context.Context) *TinyCtx {
	result := TinyCtx{}

	url, token := findURLToken(ctx)

	actMethod := strings.ToUpper(ctx.Request.Method)
	required := m.GetRequiredRole(url, actMethod)

	result.RequiredRole = required
	result.ApplicationName = m.GetServiceName()
	result.URL = url
	result.Method = actMethod
	result.SessionID = token
	result.Service = m.service

	return &result
}

func (ctx *TinyCtx) allowed() bool {
	if ctx.RequiredRole == roletype.Unknown {
		return true
	}

	return ctx.hasRole(ctx.RequiredRole)
}

func (ctx *TinyCtx) getUserKey() husk.Key {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return husk.CrazyKey()
	}

	return cookie.UserKey
}

func (ctx *TinyCtx) getUsername() string {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return "Unknown"
	}

	return cookie.Username
}

func (ctx *TinyCtx) getIP() string {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return "-1.-1.-1.-1"
	}

	return cookie.IP
}

func (ctx *TinyCtx) getLocation() string {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return "Uknown"
	}

	return cookie.Location
}

func (ctx *TinyCtx) getRole() roletype.Enum {
	result := roletype.Unknown

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

func (ctx *TinyCtx) hasRole(required roletype.Enum) bool {
	role := ctx.getRole()

	return role <= required
}

//TODO: use channels
//getAvoCookie also checks cookie validity, so repeated calls are required
func (ctx *TinyCtx) getAvoCookie() (*secure.Cookies, error) {
	if len(ctx.SessionID) == 0 {
		return nil, errors.New("SessionID empty")
	}

	result := &secure.Cookies{}
	err := mango.DoGET(result, ctx.Service.ID, "Secure.API", "login", "avo", ctx.SessionID)

	if err != nil {
		return nil, err
	}

	return result, nil
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
