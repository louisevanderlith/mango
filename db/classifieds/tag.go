package classifieds

import "github.com/louisevanderlith/mango/util"

type Tag struct {
	util.Record
	Description string
	Adverts     []*Advert `orm:"reverse(many)"`
}
