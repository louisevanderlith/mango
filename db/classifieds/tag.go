package classifieds

import (
	"github.com/louisevanderlith/husk"
)

type Tag struct {
	Description string `hsk:"size(255)"`
}

func (o Tag) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
