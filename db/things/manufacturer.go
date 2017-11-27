package things

import (
	"github.com/louisevanderlith/mango/db"
)

type Manufacturer struct {
	db.Record
	Name        string   `orm:"size(50)"`
	Description string   `orm:"null;size(255)"`
	Models      []*Model `orm:"reverse(many)"`
}