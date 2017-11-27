package comms

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/db"
)

type CommsContext struct {
	Message *db.Set
}

var Ctx *CommsContext

func NewDatabase() {
	dbName := "Communication.DB"

	registerModels()
	db.SyncDatabase(dbName)

	Ctx = &CommsContext{
		Message: db.NewSet(Message{}),
	}
}

func registerModels() {
	orm.RegisterModel(&Message{})
}
