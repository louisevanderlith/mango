package secure

import (
	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
)

type LoginTrace struct {
	db.Record
	Location string `orm:"null;size(128)"`
	IP       string `orm:"null;size(50)"`
	Allowed  bool   `orm:"default(true)"`
	User     *User  `orm:"rel(fk)"`
}

func (o LoginTrace) Validate() (bool, error) {
	return util.ValidateStruct(o)
}