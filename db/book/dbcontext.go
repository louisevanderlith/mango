package book

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	Vehicle *db.Set
}

var Ctx *Context

func NewDatabase() {
	dbName := "Book.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		registerModels()
		db.SyncDatabase(dbSource)

		Ctx = &Context{
			Vehicle: db.NewSet(Vehicle{}),
		}
	}
}

func registerModels() {
	orm.RegisterModel(new(Vehicle))
}
