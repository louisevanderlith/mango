package things

import (
	"github.com/louisevanderlith/mango/db"
)

type Model struct {
	db.Record
	Manufacturer *Manufacturer `orm:"rel(fk)"`
	Name           string `orm:"size(50)"`
}