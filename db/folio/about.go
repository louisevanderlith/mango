package folio

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type About struct {
	db.Record
	SectionText string   `orm:"size(256)"`
	Profile     *Profile `orm:"rel(fk)" json:",omitempty"`
}

func (o About) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
