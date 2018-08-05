package things

import "github.com/louisevanderlith/husk"

type Category struct {
	Name          string `hsk:"size(50)"`
	Description   string `hsk:"size(255)"`
	SubCategories Subcategories
}

func (o Category) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
