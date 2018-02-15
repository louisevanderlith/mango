package book

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/db"
)

type Context struct {
	Vehicle *db.Set
}

var Ctx *Context

func NewDatabase() {
	dbName := "Book.DB"

	registerModels()
	db.SyncDatabase(dbName)

	Ctx = &Context{
		Vehicle: db.NewSet(Vehicle{}),
	}
}

func registerModels() {
	orm.RegisterModel(new(Vehicle))
}
