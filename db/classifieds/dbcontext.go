package classifieds

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/util"
)

func NewDatabase() {
	dbName := "Classifieds.DB"
	util.BuildDatabase(registerModels, dbName)
}

func registerModels() {
	orm.RegisterModel(new(Advert), new(CarAdvert), new(Tag))
}
