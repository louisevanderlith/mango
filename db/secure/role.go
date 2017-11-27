package secure

import (
	"github.com/louisevanderlith/mango/util/enums"
	"github.com/louisevanderlith/mango/db"
	"github.com/astaxie/beego/orm"
)

type Role struct {
	db.Record
	User        *User          `orm:"rel(fk)"`
	Description enums.RoleType `orm:"type(int)"`
}

func (obj *User) LoadRoles() error{
	o := orm.NewOrm()
	_, err := o.LoadRelated(&obj, "role")

	return err
}

func GetRolesTypes(roles []*Role) []enums.RoleType{
	var result []enums.RoleType

	for _, v := range roles {
		result = append(result, v.Description)
	}

	return result
}