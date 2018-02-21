package secure

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

type Role struct {
	db.Record
	User        *User          `orm:"rel(fk)"`
	Description enums.RoleType `orm:"type(int)"`
}

func (o Role) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}

func (obj *User) LoadRoles() error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(&obj, "role")

	return err
}

func GetRolesTypes(roles []*Role) []enums.RoleType {
	var result []enums.RoleType

	for _, v := range roles {
		result = append(result, v.Description)
	}

	return result
}
