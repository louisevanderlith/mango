package artifact

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	Uploads db.Setter
	BLOBs   db.Setter
}

var Ctx *Context

func NewDatabase() {
	dbName := "Artifact.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		Ctx = &Context{
			Uploads: db.NewDBSet(Upload{}),
			BLOBs:   db.NewDBSet(Blob{}),
		}

		err = db.SyncDatabase(Ctx, dbSource)

		if err != nil {
			panic(err)
		}
	}
}
