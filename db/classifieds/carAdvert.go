package classifieds

import (
	"time"

	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db"
)

type CarAdvert struct {
	util.BaseRecord
	Make          string    `orm:"size(50)"`
	Model         string    `orm:"size(50)"`
	Info          string    `orm:"size(128)"`
	Year          int       `orm:"null"`
	Odometer      int       `orm:"null"`
	HasPapers     bool      `orm:"default(false)"`
	LicenseExpiry time.Time `orm:"type(date)"`
}

func (obj *CarAdvert) Insert() (int64, error) {
	return db.Insert(obj)
}

func (obj *CarAdvert) Read() error {
	return db.Read(obj)
}

func (obj *CarAdvert) ReadAll() ([]CarAdvert, error) {
	var data []CarAdvert
	_, err := db.ReadAll(obj, data)
	return data, err
}

func (obj *CarAdvert) Update() (int64, error) {
	return db.Update(obj)
}

func (obj *CarAdvert) Delete() error {
	obj.Deleted = true
	_, err := db.Update(obj)

	return err
}
