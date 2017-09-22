package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astaxie/beego/orm"
)

func BuildDatabase(regModels func(), instanceKey, dbName, discoveryURL string) {
	regModels()

	name := "default"
	dbPath, err := getConnectionString(instanceKey, dbName, discoveryURL)

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

func getConnectionString(instanceKey string, dbName string, discoveryURL string) (string, error) {
	var result string
	var err error

	discoveryRoute := fmt.Sprintf("%s%s/%s", discoveryURL, instanceKey, dbName)
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
			msg := fmt.Sprintf("Couldn't find a application for %s", dbName)
			err = errors.New(msg)
		}
	}

	log.Print(result)
	return result, err
}
