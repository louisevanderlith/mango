package secure

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	LoginTrace *db.Set
	Role       *db.Set
	User       *db.Set
}

var Ctx *Context

func NewDatabase() {
	dbName := "Secure.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		registerModels()
		db.SyncDatabase(dbSource)

		Ctx = &Context{
			LoginTrace: db.NewSet(LoginTrace{}),
			Role:       db.NewSet(Role{}),
			User:       db.NewSet(User{}),
		}
	}
}

func registerModels() {
	orm.RegisterModel(new(User), new(LoginTrace), new(Role))
}
