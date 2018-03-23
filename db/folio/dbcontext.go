package folio

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	About      *db.Set
	Portfolio  *db.Set
	Profile    *db.Set
	SocialLink *db.Set
	Header     *db.Set
}

var Ctx *Context

func NewDatabase() {
	dbName := "Folio.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		registerModels()
		db.SyncDatabase(dbSource)

		Ctx = &Context{
			About:      db.NewSet(About{}),
			Portfolio:  db.NewSet(Portfolio{}),
			Profile:    db.NewSet(Profile{}),
			SocialLink: db.NewSet(SocialLink{}),
			Header:     db.NewSet(Header{}),
		}
	}
}

func registerModels() {
	orm.RegisterModel(new(About), new(Portfolio), new(Profile), new(SocialLink), new(Header))
}
