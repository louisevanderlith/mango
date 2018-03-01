package classifieds

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	Advert    *db.Set
	CarAdvert *db.Set
	Tag       *db.Set
}

var Ctx *Context

func NewDatabase() {
	dbName := "Classifieds.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		registerModels()
		db.SyncDatabase(dbSource)

		Ctx = &Context{
			Advert:    db.NewSet(Advert{}),
			CarAdvert: db.NewSet(CarAdvert{}),
			Tag:       db.NewSet(Tag{}),
		}
	}
}

func registerModels() {
	orm.RegisterModel(new(Advert), new(CarAdvert), new(Tag))
}
