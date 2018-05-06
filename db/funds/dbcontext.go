package funds

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	AvoCredit   *db.Set
	LineItem    *db.Set
	Requisition *db.Set
	Transaction *db.Set
}

var Ctx *Context

func NewDatabase() {
	dbName := "Funds.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		registerModels()
		db.SyncDatabase(dbSource)

		Ctx = &Context{
			AvoCredit:   db.NewSet(AvoCredit{}),
			LineItem:    db.NewSet(LineItem{}),
			Requisition: db.NewSet(Requisition{}),
			Transaction: db.NewSet(Transaction{}),
		}
	}
}

func registerModels() {
	orm.RegisterModel(new(AvoCredit), new(LineItem), new(Requisition), new(Transaction))
}
