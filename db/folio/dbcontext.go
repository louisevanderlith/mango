package folio

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Context struct {
	Abouts      db.Setter
	Portfolios  db.Setter
	Profiles    db.Setter
	SocialLinks db.Setter
	Headers     db.Setter
}

var Ctx *Context

func NewDatabase() {
	dbName := "Folio.DB"
	dbSource, err := util.GetServiceURL(dbName, false)

	if err == nil {
		Ctx = &Context{
			Abouts:      db.NewDBSet(About{}),
			Portfolios:  db.NewDBSet(Portfolio{}),
			Profiles:    db.NewDBSet(Profile{}),
			SocialLinks: db.NewDBSet(SocialLink{}),
			Headers:     db.NewDBSet(Header{}),
		}

		err = db.SyncDatabase(Ctx, dbSource)

		if err != nil {
			panic(err)
		}
	}
}
