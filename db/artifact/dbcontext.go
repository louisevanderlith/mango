package artifact

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	Upload *db.Set
	BLOB   *db.Set
}

var Ctx *Context

func NewDatabase() {
	dbName := "Artifact.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		registerModels()
		db.SyncDatabase(dbSource)

		Ctx = &Context{
			Upload: db.NewSet(Upload{}),
			BLOB:   db.NewSet(Blob{}),
		}
	}
}

func registerModels() {
	orm.RegisterModel(new(Upload), new(Blob))
}
