package util

import (
	"log"

	"github.com/astaxie/beego/orm"
)

func BuildDatabase(regModels func(), instanceKey, dbName, discoveryURL string) {
	regModels()

	name := "default"
	dbPath, err := GetServiceURL(instanceKey, dbName, discoveryURL)

	if err != nil {
		log.Print(err)
	} else {
		driverName := "postgres"
		err := orm.RegisterDataBase(name, driverName, dbPath)

		if err != nil {
			log.Print("Please ensure that you have created your Database.")
		} else {
			orm.RunSyncdb(name, false, false)
		}
	}
}
