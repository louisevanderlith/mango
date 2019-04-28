package control

import (
	"errors"
	"fmt"
	"strings"

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

func NewTinyCtx(m *ControllerMap, method, url, token string) *TinyCtx {
	result := TinyCtx{}

	actMethod := strings.ToUpper(method)
	required := m.GetRequiredRole(url, actMethod)

	result.RequiredRole = required
	result.ApplicationName = m.GetServiceName()
	result.URL = url
	result.Method = actMethod
	result.SessionID = token
	result.Service = m.service

	return &result
}

func (ctx *TinyCtx) allowed() (bool, error) {
	if ctx.RequiredRole == roletype.Unknown {
		return true, nil
	}

	return ctx.hasRole(ctx.RequiredRole)
}

func (ctx *TinyCtx) hasRole(required roletype.Enum) (bool, error) {
	role, err := ctx.getRole()

	if err != nil {
		return false, err
	}

	return role <= required, nil
}

func (ctx *TinyCtx) getRole() (roletype.Enum, error) {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return roletype.Unknown, err
	}

	appName := ctx.ApplicationName

	role, ok := cookie.UserRoles[appName]

	if !ok {
		msg := fmt.Errorf("application permission required. %s", appName)
		return roletype.Unknown, msg
	}

	return role, nil
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
		return "Unknown"
	}

	return cookie.Location
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

	if idx == -1 {
		return url, ""
	}

	tokenIdx := strings.LastIndex(url, "=") + 1

	cleanURL = url[:idx]
	token = url[tokenIdx:]

	return cleanURL, token
}
