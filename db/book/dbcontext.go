package book

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	Vehicles     db.Setter
	VINs         db.Setter
	Services     db.Setter
	ServiceItems db.Setter
}

var Ctx *Context

func NewDatabase() {
	dbName := "Book.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		Ctx = &Context{
			Vehicles: db.NewDBSet(Vehicle{}),
			VINs:     db.NewDBSet}

		err = db.SyncDatabase(Ctx, dbSource)

		if err != nil {
			panic(err)
		}
	}
}
