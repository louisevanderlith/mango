package classifieds

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/db"
)

type ClassifiedsContext struct {
	Advert *db.Set
	CarAdvert *db.Set
	Tag *db.Set
}

var Ctx *ClassifiedsContext

func NewDatabase() {
	dbName := "Classifieds.DB"

	registerModels()
	db.SyncDatabase(dbName)

	Ctx = &ClassifiedsContext{
		Advert: db.NewSet(Advert{}),
		CarAdvert: db.NewSet(CarAdvert{}),
		Tag: db.NewSet(Tag{}),
	}
}

func registerModels() {
	orm.RegisterModel(new(Advert), new(CarAdvert), new(Tag))
}
