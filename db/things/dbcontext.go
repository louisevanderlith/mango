package things

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/db"
)

type ThingsContext struct {
	Category *db.Set
	Manufacturer *db.Set
	Model *db.Set
	SubCategory *db.Set
}

var Ctx *ThingsContext

func NewDatabase() {
	dbName := "Things.DB"

	registerModels()
	db.SyncDatabase(dbName)

	Ctx = &ThingsContext{
		Category: db.NewSet(Category{}),
		Manufacturer: db.NewSet(Manufacturer{}),
		Model: db.NewSet(Model{}),
		SubCategory: db.NewSet(SubCategory{}),
	}
}

func registerModels() {
	orm.RegisterModel(new(Category), new(SubCategory), new(Manufacturer), new(Model))
}
