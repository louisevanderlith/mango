package secure

import (
	"log"

	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

type Role struct {
	util.Record
	User        *User          `orm:"rel(fk)"`
	Description enums.RoleType `orm:"type(int)"`
}

func addUserRole(user User) {
	role := Role{
		User:        &user,
		Description: enums.User}

	o := orm.NewOrm()
	_, err := o.Insert(&role)

	if err != nil {
		log.Print(err)
	}
}
