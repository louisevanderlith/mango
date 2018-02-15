package book

import (
	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
)

type Service struct {
	db.Record
	Vehicle      *Vehicle
	DoneBy       string
	Odometer     int64
	LicensePlate string
	Items        []*ServiceItem
}

func (o Service) Validate() (bool, error) {
	return util.ValidateStruct(o)
}
