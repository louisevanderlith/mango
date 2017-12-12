package comment

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/db"
)

type Context struct {
	Comment *db.Set
}

var Ctx *Context

func NewDatabase() {
	dbName := "Comment.DB"

	registerModels()
	db.SyncDatabase(dbName)

	Ctx = &Context{
		Comment: db.NewSet(Comment{}),
	}
}

func registerModels() {
	orm.RegisterModel(new(Comment))
}
