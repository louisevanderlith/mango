package secure

import (
	"errors"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/husk"

	"strings"

	"github.com/louisevanderlith/mango/util/enums"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name          string `hsk:"size(75)"`
	Verified      bool   `hsk:"default(false)"`
	Email         string `hsk:"size(128)"`
	ContactNumber string `hsk:"size(20)"`
	Password      string
	LoginDate     time.Time
	LoginTraces   LoginTraces
	Roles         Roles
}

const cost int = 11

func (user User) Valid() (bool, error) {
	var issues []string

	if len(user.Password) < 6 {
		issues = append(issues, "Password must be atleast 6 characters.")
	}

	valid, common := husk.ValidateStruct(&user)

	if !valid {
		issues = append(issues, common.Error())
	}

	isValid := len(issues) < 1
	finErr := errors.New(strings.Join(issues, "\r\n"))

	return isValid, finErr
}

func (user User) Exists() (bool, error) {
	cond := orm.NewCondition()
	filter := cond.Or("Email", user.Email).Or("ContactNumber", user.ContactNumber)

	o := orm.NewOrm()
	result := o.QueryTable("user").SetCond(filter).Exist()

	var err error

	if !result {
		err = errors.New("User already exists.")
	}

	return result, err
}

// Login will attempt to authenticate a user
func Login(identifier string, password []byte, ip string, location string) (passed bool, userID int64, roles []enums.RoleType) {

	if identifier != "" && len(password) > 0 {
		userRec := getUser(identifier)
		user := userRec.Data().(*User)

		if user != nil {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), password)
			passed = err == nil

			if !passed {
				log.Print("Login: ", err)
			} else {
				roles = GetRolesTypes(user.Roles)
			}

			trace := LoginTrace{
				Allowed:  passed,
				Location: location,
				IP:       ip,
			}

			user.LoginTraces = append(user.LoginTraces, &trace)

			ctx.Users.Update(userRec)

			if err != nil {
				log.Print("Login: ", err)
			}
		}
	}

	return passed, userID, roles
}

func (user *User) SecurePassword() {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)

	if err != nil {
		log.Print("securePassword: ", err)
	}

	user.Password = string(hashedPwd)
}

func getUser(identifier string) userRecord {
	record := ctx.Users.FindFirst(func(o husk.Dataer) bool {
		obj := o.(*User)
		return obj.Email == identifier || obj.ContactNumber == identifier
	})

	return record
}

func GetUsers() []userRecord {
	result := ctx.Users.Find(1, 10, func(o husk.Dataer) bool {
		return true
	})

	return result
}
