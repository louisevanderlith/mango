package secure

import (
	"errors"
	"log"
	"time"

	"github.com/louisevanderlith/db"

	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
	"golang.org/x/crypto/bcrypt"
)

// User database model
type User struct {
	db.Record
	Name          string `orm:"size(75)"`
	Verified      bool   `orm:"default(false)"`
	Email         string `orm:"size(128)"`
	ContactNumber string `orm:"size(20)"`
	Password      string
	LoginDate     time.Time   `orm:"auto_now_add"`
	LoginTraces   LoginTraces `orm:"reverse(many)"`
	Roles         Roles       `orm:"reverse(many)"`
}

var cost int

func init() {
	cost = 11
}

func (user User) Validate() (bool, error) {
	var issues []string

	if len(user.Password) < 6 {
		issues = append(issues, "Password must be atleast 6 characters.")
	}

	valid, common := util.ValidateStruct(&user)

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
		user := getUser(identifier)

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
				User:     user}

			_, err = Ctx.LoginTraces.Create(&trace)
			userID = user.Id

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

func correctIdentifier(identifier string) User {
	var result User

	if identifier != "" {
		if _, err := strconv.ParseInt(identifier, 10, 64); err == nil {
			result = User{
				ContactNumber: identifier,
			}
		} else {
			result = User{
				Email: identifier,
			}
		}
	}

	return result
}

func getUser(identifier string) *User {
	var result *User

	filter := correctIdentifier(identifier)
	record, err := Ctx.Users.ReadOne(&filter, "Roles")

	if record != nil && err == nil {
		result = record.(*User)
	}

	return result
}

func GetUsers() (result Users, err error) {
	err = Ctx.Users.Read(&User{}, &result)

	return result, err
}
