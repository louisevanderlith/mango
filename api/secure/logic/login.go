package logic

import (
	"encoding/json"

	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/mango/core/secure"
	"github.com/louisevanderlith/mango/util/control"
	uuid "github.com/nu7hatch/gouuid"
)

// AttemptLogin returns SessionID, if error is not nil
func AttemptLogin(ctx *context.Context) (string, error) {

	authReq := secure.Authentication{}
	err := json.Unmarshal(ctx.Input.RequestBody, &authReq)

	if err != nil {
		return "", err
	}

	cooki, err := secure.Login(authReq)

	if err != nil {
		return "", err
	}

	u4, _ := uuid.NewV4()
	sessionID := u4.String()

	control.CreateAvo(ctx, cooki, sessionID)

	return sessionID, nil
}
