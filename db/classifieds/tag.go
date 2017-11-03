package classifieds

import "github.com/louisevanderlith/mango/util"

type Tag struct {
	util.BaseRecord
	Description string
	Adverts     []*Advert `orm:"reverse(many)"`
}

func (obj *Tag) Insert() (int64, error) {
	return db.Insert(obj)
}

func (obj *Tag) Read() error {
	return db.Read(*obj)
}

func (obj *Tag) Update() (int64, error) {
	return db.Update(obj)
}

func (obj *Tag) Delete() error {
	_, err := db.Delete(obj)

	return err
}