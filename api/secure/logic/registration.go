package logic

import "errors"
import "github.com/louisevanderlith/mango/db/secure"

type Registration struct {
	Name           string
	Email          string
	ContactNumber  string
	Password       string
	PasswordRepeat string
}

func SaveRegistration(r Registration) error {
	var err error

	if r.Password == r.PasswordRepeat {
		user := secure.User{
			Name:          r.Name,
			Email:         r.Email,
			ContactNumber: r.ContactNumber,
			Password:      r.Password}

		err = secure.CreateUser(user)
	} else {
		err = errors.New("Passwords don't match")
	}

	return err
}
