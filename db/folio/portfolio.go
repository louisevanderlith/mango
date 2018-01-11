package folio

import (
	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
)

type Portfolio struct {
	db.Record
	ImageID int64
	URL     string   `orm:"size(128)"`
	Name    string   `orm:"size(50)"`
	Profile *Profile `orm:"rel(fk)"`
}

func (o Portfolio) Validate() (bool, error) {
	return util.ValidateStruct(o)
}
