package secure

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/util"
)

func NewDatabase(instanceKey, discoveryURL string) {
	dbName := "Secure.DB"
	util.BuildDatabase(registerModels, instanceKey, dbName, discoveryURL)
}

func registerModels() {
	orm.RegisterModel(new(User), new(LoginTrace))
}
