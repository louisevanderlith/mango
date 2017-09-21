package classifieds

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	registerModels()
	buildDatabase()
}

func buildDatabase() {
	name := "default"
	dbPath := beego.AppConfig.String("dbPath")
	driverName := "postgres"

	err := orm.RegisterDataBase(name, driverName, dbPath)

	if err != nil {
		panic("Please ensure that you have created your Database.")
	} else {
		orm.RunSyncdb(name, false, false)
	}
}

func registerModels() {
	orm.RegisterModel(
		new(Advert),
		new(CarAdvert),
		new(Comment),
		new(LoginTrace),
		new(Profile),
		new(Tag),
		new(Upload),
		new(User))
}
