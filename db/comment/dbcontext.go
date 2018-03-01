package comment

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	Comment *db.Set
}

var Ctx *Context

func NewDatabase() {
	dbName := "Comment.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		registerModels()
		db.SyncDatabase(dbSource)

		Ctx = &Context{
			Comment: db.NewSet(Comment{}),
		}
	}
}

func registerModels() {
	orm.RegisterModel(new(Comment))
}
