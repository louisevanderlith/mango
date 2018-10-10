package secure

import (
	"errors"
	"strings"

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

	user := userRec.Data().(*User)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authReq.Password))

	if err != nil {
		return nil, err
	}

	passed := err == nil
	user.AddTrace(getLoginTrace(authReq, passed))
	err = ctx.Users.Update(userRec)

	if err != nil {
		return nil, err
	}

	defer ctx.Users.Save()

	if !passed {
		return nil, errors.New("login failed")
	}

	return control.NewCookies(userRec.GetKey(), user.Name, ip, location, user.RoleMap()), nil
}
