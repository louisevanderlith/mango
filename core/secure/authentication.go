package secure

import (
	"errors"
	"strings"

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

	if len(authReq.Password) < 6 {
		return nil, errors.New("password must be longer than 6 characters")
	}

	if !strings.Contains(authReq.Email, "@") {
		return nil, errors.New("email is invalid")
	}

	userRec := getUser(authReq.Email)

	if userRec == nil {
		return nil, errors.New("user not found")
	}

	ctx.Users.Update(userRec)
	defer ctx.Users.Save()

	user := userRec.Data().(*User)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authReq.Password))

	if err != nil {
		return nil, err
	}

	passed = err == nil
	userKey = userRec.GetKey()
	username = user.Name

	user.AddTrace(getLoginTrace(authReq, passed))

	if !passed {
		return nil, errors.New("login failed")
	}

	return control.NewCookies(userKey, username, ip, location), nil
}
