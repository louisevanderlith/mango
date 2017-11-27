package classifieds

import (
	"time"

	"github.com/louisevanderlith/mango/db"
)

type Advert struct {
	db.Record
	UserID     int64
	DateListed time.Time `orm:"type(datetime)"`
	Price      int
	Negotiable bool
	Tags       []*Tag    `orm:"rel(m2m)"`
	Location   string    `orm:"size(128)"`
}
