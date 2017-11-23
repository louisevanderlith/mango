package things

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db"
	"github.com/astaxie/beego/orm"
)

type Manufacturer struct {
	util.BaseRecord
	Name        string   `orm:"size(50)"`
	Description string   `orm:"null;size(255)"`
	Models      []*Model `orm:"reverse(many)"`
}

func (obj *Manufacturer) Insert() (int64, error) {
	return db.Insert(obj)
}

func (obj *Manufacturer) Read() error {
	return db.Read(*obj)
}

func (obj *Manufacturer) ReadAll() ([]*Manufacturer, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("manufacturer").Filter("Deleted", false)

	var result []*Manufacturer
	_, err := qs.All(&result)

	return result, err
}

func (obj *Manufacturer) Update() (int64, error) {
	return db.Update(obj)
}

func (obj *Manufacturer) Delete() error {
	obj.Deleted = true
	_, err := db.Update(obj)

	return err
}