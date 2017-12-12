package artifact

import (
	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
)

type Upload struct {
	db.Record
	ItemID   int64
	Name     string `orm:"size(50)"`
	MimeType string `orm:"size(30)"`
	Path     string `orm:"size(255)"`
	Size     int64
}

func (o Upload) Validate() (bool, error) {
	return util.ValidateStruct(o)
}
