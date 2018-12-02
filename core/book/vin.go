package book

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/pkg/enums"
)

type VIN struct {
	Number     string
	StandardID enums.StandardType
	Vehicle    *Vehicle
}

func (o VIN) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
