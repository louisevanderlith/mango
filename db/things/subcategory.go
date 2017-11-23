package things

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db"
	"github.com/astaxie/beego/orm"
)

type SubCategory struct {
	util.BaseRecord
	Category    *Category `orm:"rel(fk)"`
	Name        string    `orm:"size(50)"`
	Description string    `orm:"size(255)"`
}

func (obj *SubCategory) Insert() (int64, error) {
	return db.Insert(obj)
}

func (obj *SubCategory) Read() error {
	return db.Read(*obj)
}

func (obj *SubCategory) ReadAll() ([]*SubCategory, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("subcategory").Filter("Deleted", false)

	var result []*SubCategory
	_, err := qs.All(&result)

	return result, err
}

func (obj *SubCategory) Update() (int64, error) {
	return db.Update(obj)
}

func (obj *SubCategory) Delete() error {
	obj.Deleted = true
	_, err := db.Update(obj)

	return err
}