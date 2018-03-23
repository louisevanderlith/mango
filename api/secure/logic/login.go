package logic

import (
	"encoding/json"

	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/mango/db/secure"
	"github.com/louisevanderlith/mango/util"
	"github.com/nu7hatch/gouuid"
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

	if util.HasAvo(sessionID) {
		passed = true
	} else {
		var l Login
		err = json.Unmarshal(ctx.Input.RequestBody, &l)

		if err == nil {
			loggedIn, userID, roles := secure.Login(l.Identifier, []byte(l.Password), l.IP, l.Location)

			if loggedIn {
				passed = true

				session := util.Cookies{
					UserID:   userID,
					IP:       l.IP,
					Location: l.Location,
					Roles:    roles}

				util.CreateAvo(ctx, session, sessionID)
			}
		}
	}

	return passed, sessionID, err
}
