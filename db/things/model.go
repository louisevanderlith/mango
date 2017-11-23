package things

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	util.BaseRecord
	Manufacturer *Manufacturer `orm:"rel(fk)"`
	Name           string `orm:"size(50)"`
}

func (obj *Model) Insert() (int64, error) {
	return db.Insert(obj)
}

func (obj *Model) Read() error {
	return db.Read(*obj)
}

func (obj *Model) ReadAll() ([]*Model, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("model").Filter("Deleted", false)

	var result []*Model
	_, err := qs.All(&result)

	return result, err
}

func (obj *Model) Update() (int64, error) {
	return db.Update(obj)
}

func (obj *Model) Delete() error {
	obj.Deleted = true
	_, err := db.Update(obj)

	return err
}