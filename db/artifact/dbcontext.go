package artifact

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/db"
)

type Context struct {
	Upload *db.Set
}

var Ctx *Context

func NewDatabase() {
	dbName := "Artifact.DB"

	registerModels()
	db.SyncDatabase(dbName)

	Ctx = &Context{
		Upload: db.NewSet(Upload{}),
	}
}

func registerModels() {
	orm.RegisterModel(new(Upload))
}
