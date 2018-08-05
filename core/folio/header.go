package folio

import "github.com/louisevanderlith/husk"

type Header struct {
	ImageID int64  `hsk:"null"`
	Heading string `hsk:"size(50)" json:",omitempty"`
	Text    string `hsk:"size(1024)" json:",omitempty"`
}

func (o Header) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
