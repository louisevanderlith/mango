package control

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/mango/util"
)

// FilterUI is used to secure web pages.
func FilterUI(service *util.Service) func(*context.Context) {
	return func(ctx *context.Context) {
		path := ctx.Input.URL()

		if strings.HasPrefix(path, "/static") || strings.Contains(path, "favicon") {
			return
		}

		tiny := newTinyCtx(service, ctx, controlMap)

		if !tiny.allowed() {
			securityURL, err := util.GetServiceURL(service.GetInstanceKey(), "Secure.API", true)

			if err == nil {
				req := ctx.Request
				moveURL := fmt.Sprintf("%s://%s%s", ctx.Input.Scheme(), req.Host, req.RequestURI)
				loginURL := buildLoginURL(securityURL, moveURL)

				ctx.Redirect(http.StatusTemporaryRedirect, loginURL)
			}
		}
	}
}

// FilterAPI is used to secure API services.
func FilterAPI(service *util.Service) func(*context.Context) {
	return func(ctx *context.Context) {
		tiny := newTinyCtx(service, ctx, controlMap)
		if !tiny.allowed() {
			ctx.Abort(http.StatusUnauthorized, "User not authorized to access this content.")
		}
	}
}

func buildLoginURL(securityURL, returnURL string) string {
	cleanReturn := removeQueries(returnURL)
	return fmt.Sprintf("%sv1/login?return=%s", securityURL, url.QueryEscape(cleanReturn))
}

func removeQueries(url string) string {
	idxOfQuery := strings.Index(url, "?")

	if idxOfQuery != -1 {
		url = url[:idxOfQuery]
	}

	return url
}
