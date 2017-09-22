package artifact

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/logic"
)

func NewDatabase(instanceKey, discoveryURL string) {
	dbName := "Artifact.DB"
	logic.BuildDatabase(registerModels, instanceKey, dbName, discoveryURL)
}

func registerModels() {
	orm.RegisterModel(new(Upload))
}
