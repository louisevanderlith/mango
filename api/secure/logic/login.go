package logic

import (
	"encoding/json"

	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/mango/core/secure"
	"github.com/louisevanderlith/mango/util/control"
	uuid "github.com/nu7hatch/gouuid"
)

type Login struct {
	Identifier string
	Password   string
	IP         string
	Location   string
}

func AttemptLogin(ctx *context.Context) (passed bool, sessionID string, err error) {
	u4, _ := uuid.NewV4()
	sessionID = u4.String()

	if control.HasAvo(sessionID) {
		passed = true
	} else {
		var l Login
		err = json.Unmarshal(ctx.Input.RequestBody, &l)

		if err == nil {
			auth := secure.Login(l.Identifier, []byte(l.Password), l.IP, l.Location)

			if auth.Passed {
				passed = true

				session := control.Cookies{
					UserID:   auth.UserID,
					Username: auth.Username,
					IP:       l.IP,
					Location: l.Location,
					Roles:    auth.Application.}

				control.CreateAvo(session, sessionID)
			}
		}
	}

	return passed, sessionID, err
}
