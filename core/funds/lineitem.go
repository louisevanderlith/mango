package funds

import (
	"github.com/louisevanderlith/mango/pkg"
)

type LineItem struct {
	Description    string
	UnitCost       int64
	UnitsRequisted int
	UnitReceived   int
}

func (o LineItem) Valid() (bool, error) {
	return util.ValidateStruct(&o)
}
