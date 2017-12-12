package things

import (
	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
)

type SubCategory struct {
	db.Record
	Category    *Category `orm:"rel(fk)"`
	Name        string    `orm:"size(50)"`
	Description string    `orm:"size(255)"`
}

func (o SubCategory) Validate() (bool, error) {
	return util.ValidateStruct(o)
}