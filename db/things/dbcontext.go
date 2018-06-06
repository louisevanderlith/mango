package things

import (
	"log"

	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	Categories    db.Setter
	Manufacturers db.Setter
	Models        db.Setter
	SubCategories db.Setter
}

var Ctx *Context

func NewDatabase() {
	dbName := "Things.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		Ctx = &Context{
			Categories:    db.NewDBSet(Category{}),
			Manufacturers: db.NewDBSet(Manufacturer{}),
			Models:        db.NewDBSet(Model{}),
			SubCategories: db.NewDBSet(Subcategory{}),
		}
		log.Print(Ctx.Categories)
		err = db.SyncDatabase(Ctx, dbSource)

		if err != nil {
			panic(err)
		}

		seedData()
	}
}

func seedData() {
	seedManufacturer()
	seedModel()
}
