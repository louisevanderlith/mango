package things

import (
	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
)

type Model struct {
	db.Record
	Name           string `orm:"size(50)"`
	Manufacturer *Manufacturer `orm:"rel(fk)"`
}

func (o Model) Validate() (bool, error) {
	return util.ValidateStruct(o)
}