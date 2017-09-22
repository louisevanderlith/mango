package classifieds

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/logic"
)

func NewDatabase(instanceKey, discoveryURL string) {
	dbName := "Classifieds.DB"
	logic.BuildDatabase(registerModels, instanceKey, dbName, discoveryURL)
}

func registerModels() {
	orm.RegisterModel(new(Advert), new(CarAdvert), new(Tag))
}
