package folio

import "github.com/louisevanderlith/husk"

type Portfolio struct {
	ImageKey husk.Key `hsk:"null"`
	URL      string   `hsk:"size(128)"`
	Name     string   `hsk:"size(50)"`
}

func (o Portfolio) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
