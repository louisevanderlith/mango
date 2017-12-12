package comms

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/db"
)

type Context struct {
	Message *db.Set
}

var Ctx *Context

func NewDatabase() {
	dbName := "Communication.DB"

	registerModels()
	db.SyncDatabase(dbName)

	Ctx = &Context{
		Message: db.NewSet(Message{}),
	}
}

func registerModels() {
	orm.RegisterModel(&Message{})
}
