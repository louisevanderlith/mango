package folio

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Header struct {
	db.Record
	ImageID int64    `orm:"null"`
	Heading string   `orm:"size(50)" json:",omitempty"`
	Text    string   `orm:"size(1024)" json:",omitempty"`
	Profile *Profile `orm:"rel(fk)" json:",omitempty"`
}

func (o Header) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
