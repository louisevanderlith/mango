package logic

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ConnectionString struct {
	Name  string
	Value string
}

type DatabaseSetting struct {
	Name         string
	Environments []ConnectionString
}

type Settings []DatabaseSetting

var settings *Settings

func init() {
	if settings == nil {
		loadSettings()
	}
}

func loadSettings() {
	content := getFileContent("./conf/database.json")

	err := json.Unmarshal(content, &settings)

	if err != nil {
		log.Fatal(err)
	}
}

func getFileContent(configPath string) []byte {
	dat, err := ioutil.ReadFile(configPath)

	if err != nil {
		panic(err)
	}

	return dat
}
