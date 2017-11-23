package things

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/util"
)

func NewDatabase() {
	dbName := "Things.DB"
	util.BuildDatabase(registerModels, dbName)
}

func registerModels() {
	orm.RegisterModel(new(Category), new(SubCategory), new(Manufacturer), new(Model))
}
