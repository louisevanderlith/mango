package control

import (
	"net/http"

	"github.com/astaxie/beego/context"
)

type rendererFunc func(ctx *context.Context)

func (f rendererFunc) Render(ctx *context.Context) {
	f(ctx)
}

func RenderUnauthorized(err error) context.Renderer {
	return rendererFunc(func(ctx *context.Context) {
		ctx.Output.Header("WWW-Authenticate", "Bearer")

		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte(err.Error()))
	})
}
