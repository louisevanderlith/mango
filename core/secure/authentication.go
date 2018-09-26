package secure

import (
	"errors"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/util/control"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	App      control.Application
	Email    string
	Password string
}

// password hashing cost
const cost int = 11

// Login will attempt to authenticate a user
func Login(authReq Authentication) (*control.Cookies, error) {
	passed := false
	userKey := husk.CrazyKey()
	username := "Unknown"
	ip := authReq.App.IP
	location := authReq.App.Location

	if len(authReq.Password) < 6 || len(authReq.Email) < 3 {
		return nil, errors.New("login details invalid")
	}

	userRec := getUser(authReq.Email)
	defer ctx.Users.Update(userRec)

	user := userRec.Data()

	if userRec.rec == nil {
		return nil, errors.New("record is nil")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authReq.Password))

	if err != nil {
		return nil, err
	}

	passed = err == nil
	userKey = userRec.rec.GetKey()
	username = user.Name

	user.AddTrace(getLoginTrace(authReq, passed))

	if !passed {
		return nil, errors.New("login failed")
	}

	return control.NewCookies(userKey, username, ip, location), nil
}
