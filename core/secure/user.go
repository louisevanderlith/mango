package secure

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/louisevanderlith/mango/util/enums"

	"github.com/louisevanderlith/husk"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name        string `hsk:"size(75)"`
	Verified    bool   `hsk:"default(false)"`
	Email       string `hsk:"size(128)"`
	Password    string `hsk:"min(6)"`
	LoginDate   time.Time
	LoginTraces []LoginTrace
	Roles       []Role
}

func (user *User) Valid() (bool, error) {
	valid, common := husk.ValidateStruct(user)

	if !valid {
		return false, common
	}

	if !strings.Contains(user.Email, "@") {
		// #falsehood
		return false, errors.New("email is invalid")
	}

	if emailExists(user.Email) {
		return false, errors.New("email already in use")
	}

	return true, nil
}

func NewUser(name, email string) (*User, error) {
	result := new(User)
	result.Name = name
	result.Email = email
	result.Verified = false

	return result, nil
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

func getUsers(page, size int) husk.Collection {
	return ctx.Users.Find(page, size, husk.Everything())
}

func getUser(email string) husk.Recorder {
	return ctx.Users.FindFirst(emailFilter(email))
}

func emailExists(email string) bool {
	return ctx.Users.Exists(emailFilter(email))
}
