package classifieds

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db"
)

type Tag struct {
	util.BaseRecord
	Description string
	Adverts     []*Advert `orm:"reverse(many)"`
}

func (obj *Tag) Insert() (int64, error) {
	return db.Insert(obj)
}

func (obj *Tag) Read() error {
	return db.Read(obj)
}

func (obj *Tag) ReadAll() ([]Tag, error) {
	var data []Tag
	_, err := db.ReadAll(obj, data)
	return data, err
}

func (obj *Tag) Update() (int64, error) {
	return db.Update(obj)
}

func (obj *Tag) Delete() error {
	obj.Deleted = true
	_, err := db.Update(obj)

	return err
}
