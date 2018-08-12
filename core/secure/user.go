package secure

import (
	"errors"
	"log"
	"time"

	"github.com/louisevanderlith/mango/util/enums"

	"github.com/louisevanderlith/husk"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name        string `hsk:"size(75)"`
	Verified    bool   `hsk:"default(false)"`
	Email       string `hsk:"size(128)"`
	Password    string `hsk:"min(6)`
	LoginDate   time.Time
	LoginTraces []LoginTrace
	Roles       []Role
}

const cost int = 11

func (user User) Valid() (bool, error) {

	if len(user.Password) < 6 {
		return false, errors.New("passwor must be atleast 6 characters")
	}

	valid, common := husk.ValidateStruct(&user)

	if !valid {
		return false, common
	}

	if emailExists(user.Email) {
		return false, errors.New("email already in use")
	}

	return true, nil
}

// Login will attempt to authenticate a user
func Login(authReq AuthRequest) *AuthResponse {
	passed := false
	userID := int64(-1)
	username := "Unknown"
	app := authReq.GetApplication()

	if len(authReq.Password) == 0 || len(authReq.Email) < 3 {
		return NewAuthResponse(passed, userID, username, app)
	}

	userRec := getUser(authReq.Email)
	defer ctx.Users.Update(userRec)

	user := userRec.Data()

	if userRec.rec == nil {
		return NewAuthResponse(passed, userID, username, app)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), authReq.Password)
	passed = err == nil
	userID = userRec.rec.GetID()
	username = user.Name

	if !passed {
		return NewAuthResponse(passed, userID, username, app)
	}

	app.SetRole()

	trace := LoginTrace{
		Allowed:         passed,
		Location:        authReq.Location,
		IP:              authReq.IP,
		ApplicationName: authReq.ApplicationName,
		InstanceID:      authReq.InstanceID,
	}

	user.LoginTraces = append(user.LoginTraces, trace)

	if err != nil {
		log.Print("Login: ", err)
	}

	return NewAuthResponse(passed, userID, username, app)
}

func (user *User) SecurePassword() {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)

	if err != nil {
		log.Print("securePassword: ", err)
	}

	user.Password = string(hashedPwd)
}

func (user *User) AddRole(appName string, role enums.RoleType) {
	appRole := Role{appName, role}
	user.Roles = append(user.Roles, appRole)
}

func GetUsers(page, pageSize int) userSet {
	result, _ := ctx.Users.Find(page, pageSize, func(o User) bool {
		return true
	})

	return result
}

func getUser(email string) userRecord {
	record, _ := ctx.Users.FindFirst(func(o User) bool {
		return o.Email == email
	})

	return record
}

func emailExists(email string) bool {
	exists, err := ctx.Users.Exists(func(obj User) bool {
		return obj.Email == email
	})

	if err != nil {
		return true
	}

	return exists
}
