package secure

import (
	"log"

	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/util"
)

type LoginTrace struct {
	util.Record
	Location string `orm:"null;size(128)"`
	IP       string `orm:"null;size(50)"`
	Allowed  bool   `orm:"default(true)"`
	User     *User  `orm:"rel(fk)"`
}

func createLoginTrace(l LoginTrace) error {
	o := orm.NewOrm()
	_, err := o.Insert(&l)

	if err != nil {
		log.Print(err)
	}

	return err
}
