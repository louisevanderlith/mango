package control

import (
	"github.com/astaxie/beego"
)

type IDController struct {
	beego.Controller
	tinyCtx *tinyCtx
}

func (req *IDController) Prepare() {
	req.tinyCtx = newTinyCtx(req.Ctx)
}
