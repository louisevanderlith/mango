package logic

import (
	"encoding/json"
	"log"

	"github.com/louisevanderlith/mango/util"
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

func loadSettings() *Settings {
	dbConfPath := util.FindFilePath("database.json", "conf")
	content := util.GetFileContent(dbConfPath)

	var settings *Settings
	err := json.Unmarshal(content, &settings)

	if err != nil {
		log.Printf("loadSettings: ", err)
	}

	return settings
}
