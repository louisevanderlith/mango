package logic

import (
	"encoding/json"
	"errors"

	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/mango/core/secure"
	"github.com/louisevanderlith/mango/util/control"
	uuid "github.com/nu7hatch/gouuid"
)

func AttemptLogin(ctx *context.Context) (passed bool, sessionID string, err error) {
	u4, _ := uuid.NewV4()
	sessionID = u4.String()

	if control.HasAvo(sessionID) {
		return true, sessionID, nil
	}

	var authReq secure.Authentication
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

	control.CreateAvo(ctx, auth, sessionID)

	return passed, sessionID, err
}
