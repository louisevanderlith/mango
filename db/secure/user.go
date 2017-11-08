package secure

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/louisevanderlith/mango/db"

	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/util"
	"golang.org/x/crypto/bcrypt"
)

// User database model
type User struct {
	util.BaseRecord
	Name          string `orm:"size(75)"`
	Verified      bool   `orm:"default(false)"`
	Email         string `orm:"size(128)"`
	ContactNumber string `orm:"size(20)"`
	Password      string
	LoginDate     time.Time     `orm:"auto_now_add;type(datetime)"`
	LoginTraces   []*LoginTrace `orm:"reverse(many)"`
	Roles         []*Role       `orm:"reverse(many)"`
}

var cost int

func init() {
	cost = 10
}

// CreateUser will create a new user
func CreateUser(user User) error {
	err := validateUser(user)

	if err == nil && !exists(user) {
		securePassword(&user)

		o := orm.NewOrm()
		_, err = o.Insert(&user)

		if err == nil {
			addUserRole(user)
		}
	}

	return err
}

func validateUser(user User) error {
	var err error

	if user.Name == "" {
		err = errors.New("Name is invalid")
	}

	if user.Email == "" {
		err = errors.New("Email is invalid")
	}

	if user.ContactNumber == "" {
		err = errors.New("Contact Number is invalid")
	}

	if len(user.Password) < 6 {
		err = errors.New("Password must be atleast 6 characters")
	}

	return err
}

// Login will attempt to authenticate a user
func Login(identifier string, password []byte, ip string, location string) (bool, int64) {
	var passed bool
	var userID int64

	if identifier != "" && len(password) > 0 {
		user := getUser(identifier)

		if user != nil {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), password)
			passed = err == nil

			if !passed {
				log.Print(err)
			}

			trace := LoginTrace{
				Allowed:  passed,
				Location: location,
				IP:       ip,
				User:     user}

			err = createLoginTrace(trace)
			userID = user.ID

			if err != nil {
				log.Print(err)
			}
		}
	}

	return passed, userID
}

func securePassword(user *User) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)

	if err != nil {
		log.Print(err)
	}

	user.Password = string(hashedPwd)
}

func getUserByID(userID int64) *User {
	o := orm.NewOrm()
	user := new(User)
	user.ID = userID

	err := o.Read(&user)

	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		msg := fmt.Sprintf("Couldn't find user with ID %v", userID)
		log.Print(msg)
	}

	return user
}

func getUser(identifier string) *User {
	var result User
	o := orm.NewOrm()

	cond := orm.NewCondition()
	filter := cond.Or("Email", identifier).Or("ContactNumber", identifier)

	err := o.QueryTable("user").SetCond(filter).One(&result)

	if err == orm.ErrNoRows {
		msg := fmt.Sprintf("Couldn't find user with identifier %s", identifier)
		log.Print(msg)
	}

	return &result
}

func exists(user User) bool {
	o := orm.NewOrm()

	cond := orm.NewCondition()
	filter := cond.Or("Email", user.Email).Or("ContactNumber", user.ContactNumber)

	result := o.QueryTable("user").SetCond(filter).Exist()

	return result
}

func dropUser(user User) error {
	o := orm.NewOrm()

	_, err := o.Delete(&user)

	if err != nil {
		log.Print(err)
	}

	return err
}

func (obj *User) Insert() (int64, error) {
	return db.Insert(obj)
}

func (obj *User) Read() error {
	return db.Read(*obj)
}

func (obj *User) Update() (int64, error) {
	return db.Update(obj)
}

func (obj *User) Delete() error {
	obj.Deleted = true
	_, err := db.Update(obj)

	return err
}
