package things

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/util"
)

func NewDatabase(instanceKey, discoveryURL string) {
	dbName := "Things.DB"
	util.BuildDatabase(registerModels, instanceKey, dbName, discoveryURL)
}

func registerModels() {
	orm.RegisterModel(new(Category), new(Manufacturer), new(Model))
}
