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

func (user User) Valid() (bool, error) {

	if len(user.Password) < 6 {
		return false, errors.New("password must be atleast 6 characters")
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

func NewUser(name, email string) *User {
	result := new(User)
	result.Name = name
	result.Email = email
	result.Verified = false

	return result
}

func (user *User) SecurePassword(plainPassword string) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(plainPassword), cost)

	if err != nil {
		log.Print("securePassword: ", err)
	}

	user.Password = string(hashedPwd)
}

func (user *User) AddRole(appName string, role enums.RoleType) {
	appRole := Role{appName, role}
	user.Roles = append(user.Roles, appRole)
}

func (user *User) AddTrace(trace LoginTrace) {
	if trace.TraceType == TraceLogin {
		user.LoginDate = time.Now()
	}

	user.LoginTraces = append(user.LoginTraces, trace)
}

func getUsers(page, size int) (userSet, error) {
	return ctx.Users.Find(page, size, func(o User) bool {
		return true
	})
}

func getUser(email string) userRecord {
	record, _ := ctx.Users.FindFirst(func(o User) bool {
		return o.Email == email
	})

	return record
}

func emailExists(email string) bool {
	return ctx.Users.Exists(func(obj User) bool {
		return obj.Email == email
	})
}
