package classifieds

import (
	"time"

	"github.com/louisevanderlith/mango/util"
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
	return db.Read(*obj)
}

func (obj *Advert) ReadAll() (*[]Advert, error) {
	return db.ReadAll(obj)
}

func (obj *Advert) Update() (int64, error) {
	return db.Update(obj)
}

func (obj *Advert) Delete() error {
	_, err := db.Delete(obj)

	return err
}
