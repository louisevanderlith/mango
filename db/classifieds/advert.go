package classifieds

import (
	"time"
)

type Advert struct {
	Record
	UserID     int64
	DateListed time.Time `orm:"type(datetime)"`
	Price      int
	Negotiable bool
	Tags       []*Tag `orm:"rel(m2m)"`
	Location   string `orm:"size(128)"`
}
