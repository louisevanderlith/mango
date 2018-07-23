package logic

import "errors"
import (
	"github.com/louisevanderlith/mango/db/secure"
	"github.com/louisevanderlith/mango/util/enums"
)

type Registration struct {
	Name           string
	Email          string
	Password       string
	PasswordRepeat string
}

func SaveRegistration(r Registration) error {
	var err error

	if r.Password == r.PasswordRepeat {

		user := &secure.User{
			Name:          r.Name,
			Email:         r.Email,
			Password:      r.Password,
		}

		userRole := secure.Role{
			User:        user,
			Description: enums.User,
		}

		user.Roles = append(user.Roles, &userRole)
		user.SecurePassword()

		_, err = secure.Ctx.Users.Create(user)
	} else {
		err = errors.New("Passwords don't match")
	}

	return err
}
