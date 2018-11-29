package logic

import (
	"net/http"

	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/control"
)

type MasterMap struct {
	*control.ControllerMap
}

func NewMasterMap(s *util.Service) *MasterMap {
	return &MasterMap{
		control.CreateControlMap(s),
	}
}

//FilterMaster is used solely by Secure.API, as it already knows all roles.
func (c *MasterMap) FilterMaster(ctx *context.Context) {
	tCtx := control.NewTinyCtx(c.ControllerMap, ctx)

	if !allowed(tCtx) {
		ctx.Abort(http.StatusUnauthorized, "Secure.API Permission Denied!")
	}
}

func allowed(t *control.TinyCtx) bool {
	//?? No idea why I'm doing this.
	if !HasAvo(t.SessionID) {
		return true
	}

	cooki := FindAvo(t.SessionID)
	role, ok := cooki.UserRoles[t.Service.Name]

	if ok {
		return role <= t.RequiredRole
	}

	return true
}
