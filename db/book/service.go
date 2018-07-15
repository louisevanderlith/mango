package book

import (
	"github.com/louisevanderlith/husk"
)

type Service struct {
	DoneBy       string
	Odometer     int64
	LicensePlate string
	Items        ServiceItems
}

func (o Service) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
