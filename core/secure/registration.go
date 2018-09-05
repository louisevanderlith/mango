package secure

import (
	"errors"

	"github.com/louisevanderlith/mango/util/control"
	"github.com/louisevanderlith/mango/util/enums"
)

type Registration struct {
	App            control.Application
	Name           string
	Email          string
	Password       string
	PasswordRepeat string
}

func Register(r Registration) (result userRecord, err error) {
	if r.Password != r.PasswordRepeat {
		err = errors.New("passwords do not match")
		return result, err
	}

	if len(r.App.Name) == 0 {
		err = errors.New("application name can not be empty")
		return result, err
	}

	if len(r.App.InstanceID) == 0 {
		err := errors.New("instance id can not be empty")
		return result, err
	}

	user := NewUser(r.Name, r.Email)

	user.SecurePassword(r.Password)
	user.AddTrace(getRegistrationTrace(r))
	user.AddRole(r.App.Name, enums.User)

	rec, err := ctx.Users.Create(user)

	return rec, err
}
