package secure

import (
	"log"

	"github.com/astaxie/beego/orm"
)

type LoginTrace struct {
	Record
	Location string `orm:"null;size(128)"`
	IP       string `orm:"null;size(50)"`
	Allowed  bool   `orm:"default(true)"`
	User     *User  `orm:"rel(fk)"`
}

func createLoginTrace(l LoginTrace) error {
	o := orm.NewOrm()
	_, err := o.Insert(&l)

	if err != nil {
		log.Panic(err)
	}

	return err
}
