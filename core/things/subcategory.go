package things

import "github.com/louisevanderlith/husk"

type Subcategory struct {
	Name        string `hsk:"size(50)"`
	Description string `hsk:"size(255)"`
}

func (o Subcategory) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
