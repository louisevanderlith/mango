package folio

import "github.com/louisevanderlith/husk"

type About struct {
	SectionText string `hsk:"size(512)"`
}

func (o About) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
