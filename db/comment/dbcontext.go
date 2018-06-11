package comment

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	Messages db.Setter
}

var Ctx *Context

func NewDatabase() {
	dbName := "Comment.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		Ctx = &Context{
			Messages: db.NewDBSet(Message{}),
		}

		err = db.SyncDatabase(Ctx, dbSource)

		if err != nil {
			panic(err)
		}
	}
}
