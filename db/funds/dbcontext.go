package funds

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	Experiences  db.Setter
	Heroes       db.Setter
	Levels       db.Setter
	LineItems    db.Setter
	Requisitions db.Setter
	Transactions db.Setter
}

var Ctx *Context

func NewDatabase() {
	dbName := "Funds.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		Ctx = &Context{
			Experiences:  db.NewDBSet(Experience{}),
			Heroes:       db.NewDBSet(Hero{}),
			Levels:       db.NewDBSet(Level{}),
			LineItems:    db.NewDBSet(LineItem{}),
			Requisitions: db.NewDBSet(Requisition{}),
			Transactions: db.NewDBSet(Transaction{}),
		}

		err = db.SyncDatabase(Ctx, dbSource)

		if err != nil {
			panic(err)
		}

		seedData()
	}
}

func seedData() {
	seedLevel()
}
