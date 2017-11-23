package secure

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db"
)

type LoginTrace struct {
	util.BaseRecord
	Location string `orm:"null;size(128)"`
	IP       string `orm:"null;size(50)"`
	Allowed  bool   `orm:"default(true)"`
	User     *User  `orm:"rel(fk)"`
}

func (obj *LoginTrace) Insert() (int64, error) {
	return db.Insert(obj)
}
