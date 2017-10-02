package classifieds

import (
	"time"

	"github.com/louisevanderlith/classifiedcore/db"
)

type Advert struct {
	db.Record
	User       *User     `orm:"rel(one)"`
	DateListed time.Time `orm:"type(datetime)"`
	Price      int
	Negotiable bool
	Tags       []*Tag     `orm:"rel(m2m)"`
	Comments   []*Comment `orm:"rel(m2m)"`
	Uploads    []*Upload  `orm:"reverse(many)"`
	Location   string     `orm:"size(128)"`
}
