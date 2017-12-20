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
	Size     int
	BLOB     *Blob  `orm:"rel(fk)"`
}

func (o Upload) Validate() (bool, error) {
	return util.ValidateStruct(o)
}
