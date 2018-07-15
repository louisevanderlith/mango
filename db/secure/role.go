package secure

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/util/enums"
)

type Role struct {
	Description enums.RoleType
}

func (o Role) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
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
