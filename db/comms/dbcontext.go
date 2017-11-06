package comms

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/util"
)

func NewDatabase() {
	dbName := "Comms.DB"
	util.BuildDatabase(registerModels, dbName)
}

func registerModels() {
	orm.RegisterModel(new(Message))
}
