package things

import (
	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
)

type Model struct {
	db.Record
	Manufacturer *Manufacturer `orm:"rel(fk)"`
	Name           string `orm:"size(50)"`
}

func (o Model) Validate() (bool, error) {
	return util.ValidateStruct(o)
}