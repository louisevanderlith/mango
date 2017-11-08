package classifieds

import (
	"time"

	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db"
)

type Advert struct {
	util.BaseRecord
	UserID     int64
	DateListed time.Time `orm:"type(datetime)"`
	Price      int
	Negotiable bool
	Tags       []*Tag `orm:"rel(m2m)"`
	Location   string `orm:"size(128)"`
}

func (obj *Advert) Insert() (int64, error) {
	return db.Insert(obj)
}

func (obj *Advert) Read() error {
	return db.Read(obj)
}

func (obj *Advert) ReadAll() ([]Advert, error) {
	var data []Advert
	_, err := db.ReadAll(obj, data)
	return data, err
}

func (obj *Advert) Update() (int64, error) {
	return db.Update(obj)
}

func (obj *Advert) Delete() error {
	obj.Deleted = true
	_, err := db.Update(obj)

	return err
}
