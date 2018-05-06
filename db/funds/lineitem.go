package funds

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type LineItem struct {
	db.Record
	Requisition    *Requisition `orm:"rel(fk)"`
	Description    string
	UnitCost       int64
	UnitsRequisted int
	UnitReceived   int
}

func (o LineItem) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
