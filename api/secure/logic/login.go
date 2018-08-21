package logic

import (
	"encoding/json"
	"errors"

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
		return true, sessionID, nil
	}

	var authReq secure.AuthRequest
	err = json.Unmarshal(ctx.Input.RequestBody, &authReq)

	if err != nil {
		return false, sessionID, err
	}

	auth := secure.Login(authReq)

	passed = auth.Passed

	if !passed {
		errMsg := errors.New("login failed")
		return passed, sessionID, errMsg
	}

	session := control.Cookies{
		UserKey:  auth.UserKey,
		Username: auth.Username,
		IP:       authReq.IP,
		Location: authReq.Location,
	}

	session.UserRoles = auth.Application.Roles

	control.CreateAvo(ctx, session, sessionID)

	return passed, sessionID, err
}
