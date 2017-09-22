package comment

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/logic"
)

func NewDatabase(instanceKey, discoveryURL string) {
	dbName := "Comment.DB"
	logic.BuildDatabase(registerModels, instanceKey, dbName, discoveryURL)
}

func registerModels() {
	orm.RegisterModel(new(Comment))
}
