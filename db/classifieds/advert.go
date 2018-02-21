package classifieds

import (
	"time"

	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
)

type Advert struct {
	db.Record
	UserID     int64
	DateListed time.Time `orm:"type(datetime)"`
	Price      int
	Negotiable bool
	Tags       []*Tag `orm:"rel(m2m)"`
	Location   string `orm:"size(128)"`
}

func (o Advert) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
