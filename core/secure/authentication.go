package secure

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/util/control"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	App      control.Application
	Email    string
	Password string
}

const cost int = 11

// Login will attempt to authenticate a user
func Login(authReq Authentication) *control.Cookies {
	passed := false
	userKey := husk.NewKey(-1)
	username := "Unknown"

	if len(authReq.Password) < 6 || len(authReq.Email) < 3 {
		return control.N NewAuthResponse(passed, userKey, username, &authReq.App)
	}

	userRec := getUser(authReq.Email)
	defer ctx.Users.Update(userRec)

	user := userRec.Data()

	if userRec.rec == nil {
		return NewAuthResponse(passed, userKey, username, &authReq.App)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authReq.Password))

	if err != nil {
		return NewAuthResponse(passed, userKey, username, &authReq.App)
	}

	passed = err == nil
	userKey = userRec.rec.GetKey()
	username = user.Name

	user.AddTrace(getLoginTrace(authReq, passed))

	if !passed {
		return NewAuthResponse(passed, userKey, username, &authReq.App)
	}

	return NewAuthResponse(passed, userKey, username, &authReq.App)
}
