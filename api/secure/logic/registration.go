package logic

import "errors"
import (
	"github.com/louisevanderlith/mango/db/secure"
	"github.com/louisevanderlith/mango/util/enums"
)

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
			Password:      r.Password,
		}

		user.Roles[0] = &secure.Role{
			User:        &user,
			Description: enums.User,
		}

		_, err = secure.Ctx.User.Create(user)
	} else {
		err = errors.New("Passwords don't match")
	}

	return err
}
