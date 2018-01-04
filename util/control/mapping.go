package control

import (
	"github.com/astaxie/beego/context"
	"strings"
	"github.com/louisevanderlith/mango/util/enums"
	"log"
	"net/http"
	"fmt"
	"net/url"
	"github.com/louisevanderlith/mango/util"
)

type MethodMap map[string]enums.RoleType
type controllerMap map[string]MethodMap

var controllerMapping controllerMap

func init() {
	controllerMapping = make(controllerMap)
}

func AddControllerMap(path string, methodMap MethodMap) {
	controllerMapping[path] = methodMap
}

func FilterUI(ctx *context.Context) {
	path := ctx.Input.URL()
	if strings.HasPrefix(path, "/static") || strings.Contains(path, "favicon") {
		return
	}

	tinyCtx := newTinyCtx(ctx)

	if !tinyCtx.allowed() {
		securityURL, err := util.GetServiceURL("Security.API", true)

		if err == nil {
			log.Println("FilterUI:", err)
			req := ctx.Request
			moveURL := fmt.Sprintf("%s://%s%s", ctx.Input.Scheme(), req.Host, req.RequestURI)
			loginURL := fmt.Sprintf("%sv1/login?return=%s", securityURL, url.QueryEscape(moveURL))

			ctx.Redirect(http.StatusTemporaryRedirect, loginURL)
		}
	}
}

func FilterAPI(ctx *context.Context) {
	tinyCtx := newTinyCtx(ctx)

	if !tinyCtx.allowed() {
		ctx.Abort(http.StatusUnauthorized, "User not authorized to access this content.")
	}
}
