package secure

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
	"github.com/louisevanderlith/mango/db"
	"github.com/astaxie/beego/orm"
)

type Role struct {
	util.BaseRecord
	User        *User          `orm:"rel(fk)"`
	Description enums.RoleType `orm:"type(int)"`
}

func (obj *Role) Insert() (int64, error) {
	return db.Insert(obj)
}

func (obj *Role) Read() error {
	return db.Read(*obj)
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