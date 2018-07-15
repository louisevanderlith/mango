package artifact

import (
	"github.com/louisevanderlith/husk"
)

type Upload struct {
	ItemID   int64
	ItemName string `hsk:"size(75)"`
	Name     string `hsk:"size(50)"`
	MimeType string `hsk:"size(30)"`
	Size     int
	BLOB     *Blob
}

func (o Upload) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
