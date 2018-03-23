package comms

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	Message *db.Set
}

var Ctx *Context

func NewDatabase() {
	dbName := "Communication.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		registerModels()
		db.SyncDatabase(dbSource)

		Ctx = &Context{
			Message: db.NewSet(Message{}),
		}
	}
}

func registerModels() {
	orm.RegisterModel(&Message{})
}
