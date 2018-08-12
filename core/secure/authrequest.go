package secure

import (
	"errors"

	"github.com/louisevanderlith/mango/util/control"
	"github.com/louisevanderlith/mango/util/enums"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	ApplicationName string
	InstanceID      uuid.UUID
	Name            string
	Email           string
	Password        string
	PasswordRepeat  string
	IP              string
	Location        string
}

const cost int = 11

// Login will attempt to authenticate a user
func Login(authReq AuthRequest) *AuthResponse {
	passed := false
	userID := int64(-1)
	username := "Unknown"
	app := authReq.GetApplication()

	if len(authReq.Password) < 6 || len(authReq.Email) < 3 {
		return NewAuthResponse(passed, userID, username, app)
	}

	userRec := getUser(authReq.Email)
	defer ctx.Users.Update(userRec)

	user := userRec.Data()

	if userRec.rec == nil {
		return NewAuthResponse(passed, userID, username, app)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authReq.Password))

	if err != nil {
		return NewAuthResponse(passed, userID, username, app)
	}

	passed = err == nil
	userID = userRec.rec.GetID()
	username = user.Name

	user.AddTrace(getLoginTrace(authReq, passed))

	if !passed {
		return NewAuthResponse(passed, userID, username, app)
	}

	return NewAuthResponse(passed, userID, username, app)
}

func (req AuthRequest) GetApplication() *control.Application {
	return control.NewApplication(req.ApplicationName, req.InstanceID)
}

func (r AuthRequest) CreateUser() (result userRecord, err error) {

	if r.Password != r.PasswordRepeat {
		err = errors.New("passwords do not match")
		return result, err
	}

	if len(r.ApplicationName) == 0 {
		err = errors.New("application name can not be empty")
		return result, err
	}

	if len(r.InstanceID) == 0 {
		err := errors.New("instance id can not be empty")
		return result, err
	}

	user := NewUser(r.Name, r.Email, false)

	user.SecurePassword(r.Password)
	user.AddTrace(getRegistrationTrace(r))
	user.AddRole(r.ApplicationName, enums.User)

	rec, err := ctx.Users.Create(user)

	return rec, err
}
