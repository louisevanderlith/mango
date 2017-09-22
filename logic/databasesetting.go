package logic

import (
	"encoding/json"
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
		loadSettings("../api/router/conf/database.json")
	}
}

func loadSettings(dbConfPath string) {
	content := getFileContent(dbConfPath)

	err := json.Unmarshal(content, &settings)

	if err != nil {
		log.Panic(err)
	}
}
