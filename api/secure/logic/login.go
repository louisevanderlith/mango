package logic

import (
	"encoding/json"

	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/mango/core/secure"
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

	sessionID := CreateAvo(ctx, cooki)

	return sessionID, nil
}
