package classifieds

import "github.com/louisevanderlith/mango/util"

type Tag struct {
	util.BaseRecord
	Description string
	Adverts     []*Advert `orm:"reverse(many)"`
}
