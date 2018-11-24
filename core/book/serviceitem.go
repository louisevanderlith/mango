package book

import (
	"github.com/louisevanderlith/husk"
)

type ServiceItem struct {
	Code        string
	Description string
}

func (o ServiceItem) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
