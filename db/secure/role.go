package secure

import "github.com/louisevanderlith/mango/util/enums"

type Role struct {
	Record
	User        *User          `orm:"rel(fk)"`
	Description enums.RoleType `orm:"type(int)"`
}
