package classifieds

import (
	"github.com/louisevanderlith/mango/db"
)

type Tag struct {
	db.Record
	Description string
	Adverts     []*Advert `orm:"reverse(many)"`
}