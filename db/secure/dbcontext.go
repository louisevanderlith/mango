package secure

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	LoginTraces db.Setter
	Roles       db.Setter
	Users       db.Setter
}

var Ctx *Context

func NewDatabase() {
	dbName := "Secure.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		Ctx = &Context{
			LoginTraces: db.NewDBSet(LoginTrace{}),
			Roles:       db.NewDBSet(Role{}),
			Users:       db.NewDBSet(User{}),
		}

		err = db.SyncDatabase(Ctx, dbSource)

		if err != nil {
			panic(err)
		}
	}
}
