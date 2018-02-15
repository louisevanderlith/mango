package common

import (
	"encoding/json"
	"strings"

	"github.com/louisevanderlith/mango/util"
)

type Region struct {
	Continent string
	Country   string
}

type Continent struct {
	CharKey
	Countries map[string][]CharKey
}

type CharKey struct {
	ID    string
	Name  string
	Start string
	End   string
}

var world []Continent

func init() {
	world = buildWorld()
}

func GetRegion(wmiCode WMICode) Region {
	var result Region

	const illegalChars = "IOQ"

	continentValid := strings.Index(illegalChars, wmiCode.ContinentCode) == -1
	countryValid := strings.Index(illegalChars, wmiCode.RegionCode) == -1

	if continentValid && countryValid {
		continent := findContinent(wmiCode.ContinentCode)

		if len(continent.Countries) > 0 {
			countries := continent.Countries[wmiCode.ContinentCode]
			county := findCountry(countries, wmiCode.RegionCode)

			result.Continent = continent.Name
			result.Country = county.Name
		}
	}

	return result
}

func findContinent(continentCode string) Continent {
	var result Continent

	for _, v := range world {
		if isBetween(continentCode, v.Start, v.End) {
			result = v
			break
		}
	}

	return result
}

func findCountry(countries []CharKey, countryCode string) CharKey {
	var result CharKey

	for _, v := range countries {
		if isBetween(countryCode, v.Start, v.End) {
			result = v
			break
		}
	}

	return result
}

func isBetween(char, start, end string) bool {
	// Check to see if a character is between to other characters in sequence.
	charValue := getValue(char)
	startValue := getValue(start)
	endValue := getValue(end)

	return startValue <= charValue && endValue >= charValue
}

func buildWorld() []Continent {
	var result []Continent

	continents, err := loadContinents()

	if err == nil {
		for _, v := range continents {
			regions, err := loadRegions(v.ID)

			continent := Continent{}
			continent.CharKey = v
			continent.Countries = make(map[string][]CharKey)

			if err == nil {
				for prefix, countries := range regions {
					continent.Countries[prefix] = countries
				}
			}

			result = append(result, continent)
		}
	}

	return result
}

func loadContinents() ([]CharKey, error) {
	worldConf := util.FindFilePath("world.json", "data")
	content := util.GetFileContent(worldConf)

	var continents []CharKey
	err := json.Unmarshal(content, &continents)

	return continents, err
}

func loadRegions(key string) (map[string][]CharKey, error) {
	regionPath := util.FindFilePath(key+".json", "data/countries")
	regionContent := util.GetFileContent(regionPath)

	var regions map[string][]CharKey
	err := json.Unmarshal(regionContent, &regions)

	return regions, err
}
