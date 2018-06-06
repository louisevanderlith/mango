package classifieds

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Tag struct {
	db.Record
	Description string  `orm:"size(255)"`
	Adverts     Adverts `orm:"reverse(many)"`
}

func (o Tag) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
