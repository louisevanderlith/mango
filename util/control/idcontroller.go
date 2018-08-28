package control

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/husk"
)

type IDController struct {
	beego.Controller
	tinyCtx *tinyCtx
}

func (req *IDController) Prepare() {
	req.tinyCtx = newTinyCtx(req.Ctx)
}

func (req *IDController) UserKey() husk.Key {
	return req.tinyCtx.getUserKey()
}

func (req *IDController) GetUsername() string {
	return req.tinyCtx.getUsername()
}
