package things

import (
	"github.com/louisevanderlith/mango/db"
)

type SubCategory struct {
	db.Record
	Category    *Category `orm:"rel(fk)"`
	Name        string    `orm:"size(50)"`
	Description string    `orm:"size(255)"`
}