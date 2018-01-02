package db

import (
	"log"
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/util"
)

func SyncDatabase(dbName string) {
	name := "default"
	dbPath, err := util.GetServiceURL(dbName, false)

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
