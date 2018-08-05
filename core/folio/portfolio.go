package folio

import "github.com/louisevanderlith/husk"

type Portfolio struct {
	ImageID int64    `hsk:"null"`
	URL     string   `hsk:"size(128)"`
	Name    string   `hsk:"size(50)"`
	Profile *Profile `json:",omitempty"`
}

func (o Portfolio) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
