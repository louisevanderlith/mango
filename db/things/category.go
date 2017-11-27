package things

import (
	"github.com/louisevanderlith/mango/db"
)

type Category struct {
	db.Record
	Name          string         `orm:"size(50)"`
	Description   string         `orm:"size(255)"`
	SubCategories []*SubCategory `orm:"reverse(many)"`
}