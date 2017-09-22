package comms

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/logic"
)

func NewDatabase(instanceKey, discoveryURL string) {
	dbName := "Comms.DB"
	logic.BuildDatabase(registerModels, instanceKey, dbName, discoveryURL)
}

func registerModels() {
	orm.RegisterModel(new(Message))
}
