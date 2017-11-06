package artifact

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/util"
)

func NewDatabase() {
	dbName := "Artifact.DB"
	util.BuildDatabase(registerModels, dbName)
}

func registerModels() {
	orm.RegisterModel(new(Upload))
}
