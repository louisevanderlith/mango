package classifieds

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	Adverts    db.Setter
	CarAdverts db.Setter
	Tags       db.Setter
}

var Ctx *Context

func NewDatabase() {
	dbName := "Classifieds.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		Ctx = &Context{
			Adverts:    db.NewDBSet(Advert{}),
			CarAdverts: db.NewDBSet(CarAdvert{}),
			Tags:       db.NewDBSet(Tag{}),
		}

		err = db.SyncDatabase(Ctx, dbSource)

		if err != nil {
			panic(err)
		}
	}
}
