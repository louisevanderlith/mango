package common

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/louisevanderlith/mango/util"
)

type Manufacturer struct {
	Category WMICategory
	Name     string
	VDSName  string // Shortname which will be used to load the correct IBrand interface
}

type ManufacturerGroup map[string]Manufacturer
type WorldGroup map[string]ManufacturerGroup

var worldGroup WorldGroup

func init() {
	worldGroup = loadManufacturerWorld()
}

func GetManufacturer(continentCode, manufacturerCode string) Manufacturer {
	var result Manufacturer

	group, ok := worldGroup[continentCode]

	if ok {
		manufacturer, hasManu := group[manufacturerCode]

		if hasManu {
			result = manufacturer
		}
	}

	return result
}

func loadManufacturerWorld() WorldGroup {
	manConf := util.FindFilePath("manufacturer.json", "data")
	content := util.GetFileContent(manConf)

	var manus []struct {
		Group string
		Code  string
		Name  string
	}

	err := json.Unmarshal(content, &manus)

	result := make(WorldGroup)

	if err == nil {
		for _, v := range manus {
			group, ok := result[v.Group]

			manu := Manufacturer{
				Category: NotSpecified,
				Name:     v.Name,
				VDSName:  getVDSName(v.Name),
			}

			if ok {
				group[v.Code] = manu
			} else {
				ngroup := make(map[string]Manufacturer)
				ngroup[v.Code] = manu
				result[v.Group] = ngroup
			}
		}
	} else {
		fmt.Print(err)
	}

	return result
}

func getVDSName(name string) string {
	var result string

	vdsNames := []string{
		"Toyota",
		"Ford",
		"Opel",
		"Mercury",
		"Mazda",
		"FIAT",
		"Mercedes",
		"Nissan",
		"Audi",
		"BMW",
		"Honda",
	}

	for _, v := range vdsNames {
		if strings.Contains(name, v) {
			result = v
			break
		}
	}

	return result
}
