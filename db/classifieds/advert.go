package classifieds

import (
	"time"

	"github.com/louisevanderlith/mango/util"
)

type Advert struct {
	util.BaseRecord
	UserID     int64
	DateListed time.Time `orm:"type(datetime)"`
	Price      int
	Negotiable bool
	Tags       []*Tag `orm:"rel(m2m)"`
	Location   string `orm:"size(128)"`
}
