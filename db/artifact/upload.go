package artifact

import (
	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
)

type Upload struct {
	util.BaseRecord
	ItemID   int64
	Name     string
	MimeType string
	Path     string
	Size     int64
}

func (obj *Upload) Insert() (int64, error) {
	return db.Insert(obj)
}

func (obj *Upload) Read() error {
	return db.Read(*obj)
}

func (obj *Upload) ReadAll() (*[]Upload, error) {
	return db.ReadAll(obj)
}

func (obj *Upload) Update() (int64, error) {
	return db.Update(obj)
}

func (obj *Upload) Delete() error {
	_, err := db.Delete(obj)

	return err
}
