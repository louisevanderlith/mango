package secure

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/util"
)

func NewDatabase() {
	dbName := "Secure.DB"
	util.BuildDatabase(registerModels, dbName)
}

func registerModels() {
	orm.RegisterModel(new(User), new(LoginTrace), new(Role))
}
