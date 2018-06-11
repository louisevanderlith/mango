package control

import (
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/mango/util/enums"
)

type tinyCtx struct {
	Controller string
	Method     string
	SessionID  string
}

const avosession = "avosession"

func newTinyCtx(ctx *context.Context) tinyCtx {
	result := tinyCtx{}

	ctrl, sess := removeToken(ctx.Request.RequestURI)

	if sess == "" {
		sess = ctx.GetCookie(avosession)
	}

	result.Controller = ctrl
	result.Method = strings.ToUpper(ctx.Request.Method)
	result.SessionID = sess

	return result
}

func (ctx tinyCtx) allowed() bool {
	result := true
	ctrlMap, hasCtrl := controllerMapping[ctx.Controller]

	if hasCtrl {
		methodMap, hasMethod := ctrlMap[ctx.Method]

		if hasMethod {
			result = hasRole(methodMap, ctx.SessionID)
		}
	}

	return result
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

func hasRole(funcRole enums.RoleType, sessionID string) bool {
	result := false

	if sessionID != "" {
		roles, _ := GetRoles(sessionID)

		if len(roles) > 0 {
			for _, val := range roles {
				if val <= funcRole {
					result = true
					break
				}
			}
		}
	}

	return result
}
