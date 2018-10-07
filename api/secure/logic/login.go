package logic

import (
	"log"
	"encoding/json"

	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/mango/core/secure"
	"github.com/louisevanderlith/mango/util/control"
)

// AttemptLogin returns SessionID, if error is not nil
func AttemptLogin(ctx *context.Context) (string, error) {
	authReq := secure.Authentication{}
	err := json.Unmarshal(ctx.Input.RequestBody, &authReq)

	if err != nil {
		return "", err
	}

	log.Printf("Authre:- %+v\n", authReq)
	cooki, err := secure.Login(authReq)

	if err != nil {
		return "", err
	}

	sessionID := control.CreateAvo(ctx, cooki)

	return sessionID, nil
}
