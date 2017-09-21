package secure

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

// User database model
type User struct {
	Record
	Profile       *Profile `orm:"rel(one)"`
	Email         string   `orm:"size(128)"`
	ContactNumber string   `orm:"size(20)"`
	Password      string
	LoginDate     time.Time     `orm:"type(date)"`
	LoginTraces   []*LoginTrace `orm:"reverse(many)"`
}

// CreateUser will create a new user
func CreateUser(user User) error {
	var err error

	if user.Email == "" {
		err = errors.New("Email is invalid")
	}

	if user.ContactNumber == "" {
		err = errors.New("Contact Number is invalid")
	}

	if err == nil && !user.exists() {
		o := orm.NewOrm()
		_, err = o.Insert(&user)
	} else {
		err = errors.New("Unable to create user. Email and Contact Number is not unique")
	}

	return err
}

func (user *User) exists() bool {
	o := orm.NewOrm()

	result := o.QueryTable("user").Filter("Email", "ContactNumber").Exist()

	return result
}
