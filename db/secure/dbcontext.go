package secure

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/db"
)

type SecureContext struct {
	LoginTrace *db.Set
	Role *db.Set
	User *db.Set
}

var Ctx *SecureContext

func NewDatabase() {
	dbName := "Secure.DB"

	registerModels()
	db.SyncDatabase(dbName)

	Ctx = &SecureContext{
		LoginTrace: db.NewSet(LoginTrace{}),
		Role: db.NewSet(Role{}),
		User: db.NewSet(User{}),
	}
}

func registerModels() {
	orm.RegisterModel(new(User), new(LoginTrace), new(Role))
}
