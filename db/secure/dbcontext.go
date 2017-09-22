package secure

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/logic"
)

func NewDatabase(instanceKey, discoveryURL string) {
	dbName := "Secure.DB"
	logic.BuildDatabase(registerModels, instanceKey, dbName, discoveryURL)
}

func registerModels() {
	orm.RegisterModel(new(User), new(LoginTrace))
}
