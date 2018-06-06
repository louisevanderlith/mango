package book

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

type VIN struct {
	db.Record
	Number     string
	StandardID enums.StandardType `orm:"type(int)"`
	Vehicle    *Vehicle
}

func (o VIN) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
