package things

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	Category     *db.Set
	Manufacturer *db.Set
	Model        *db.Set
	SubCategory  *db.Set
}

var Ctx *Context

func NewDatabase() {
	dbName := "Things.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		registerModels()
		db.SyncDatabase(dbSource)

		Ctx = &Context{
			Category:     db.NewSet(Category{}),
			Manufacturer: db.NewSet(Manufacturer{}),
			Model:        db.NewSet(Model{}),
			SubCategory:  db.NewSet(Subcategory{}),
		}

		seedData()
	}
}

func registerModels() {
	orm.RegisterModel(new(Category), new(Subcategory), new(Manufacturer), new(Model))
}

func seedData() {
	seedManufacturer()
	seedModel()
}
