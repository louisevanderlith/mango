package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/louisevanderlith/mango/db/comms"
)

func init() {
	registerModels()
}

func BuildDatabase(instanceKey string) {
	name := "default"
	dbPath, err := getConnectionString(instanceKey, "Communication.DB")

	if err != nil {
		log.Panic(err)
	} else {
		driverName := "postgres"
		err := orm.RegisterDataBase(name, driverName, dbPath)

		if err != nil {
			log.Panic("Please ensure that you have created your Database.")
		} else {
			orm.RunSyncdb(name, false, false)
		}
	}
}

func registerModels() {
	orm.RegisterModel(
		new(comms.Message))
}

func getConnectionString(appKey string, databaseName string) (string, error) {
	var result string
	var err error

	discoveryRoute := fmt.Sprintf("%s%s/%s", beego.AppConfig.String("discovery"), appKey, databaseName)
	resp, err := http.Get(discoveryRoute)
	defer resp.Body.Close()

	if err != nil {
		log.Panic(err)
	} else {
		contents, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Panic(err)
		}

		jsonErr := json.Unmarshal(contents, &result)

		if jsonErr != nil {
			log.Panic(err)
		}

		if result == "" {
			msg := fmt.Sprintf("Couldn't find a application for %s", databaseName)
			err = errors.New(msg)
		}
	}

	log.Print(result)
	return result, err
}
