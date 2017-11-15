package util

import (
	"log"

	"github.com/astaxie/beego/orm"
)

func BuildDatabase(regModels func(), dbName string) {
	regModels()

	name := "default"
	dbPath, err := GetServiceURL(dbName)

	if err != nil {
		log.Printf("BuildDatabase: ", err)
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
