package main

import (
	"encoding/json"
	"log"

	"github.com/louisevanderlith/mango/util"
)

type DomainSetting struct {
	Address string
	Name    string
}

type Settings []DomainSetting

func loadSettings() *Settings {
	dbConfPath := util.FindFilePath("domains.json")
	content := util.GetFileContent(dbConfPath)

	var settings *Settings
	err := json.Unmarshal(content, &settings)

	if err != nil {
		log.Print(err)
	}

	return settings
}
