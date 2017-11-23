package things

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db"
	"github.com/astaxie/beego/orm"
)

type Category struct {
	util.BaseRecord
	Name          string         `orm:"size(50)"`
	Description   string         `orm:"size(255)"`
	SubCategories []*SubCategory `orm:"reverse(many)"`
}

func (obj *Category) Insert() (int64, error) {
	return db.Insert(obj)
}

func (obj *Category) Read() error {
	return db.Read(*obj)
}

func (obj *Category) ReadAll() ([]*Category, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("category").Filter("Deleted", false)

	var result []*Category
	_, err := qs.All(&result)

	return result, err
}

func (obj *Category) Update() (int64, error) {
	return db.Update(obj)
}

func (obj *Category) Delete() error {
	obj.Deleted = true
	_, err := db.Update(obj)

	return err
}
