package brands

import (
	"fmt"
	"strings"

	"github.com/louisevanderlith/mango/util/vin/common"
)

type Data struct {
	Model        string
	BodyStyle    string
	Doors        string
	EngineModel  string
	EngineSize   string
	Trim         string
	Transmission string
	Gears        string
	Extra        string
}

type Dirty struct {
	Models        []string
	BodyStyles    []string
	Doors         []string
	EngineModels  []string
	EngineSizes   []string
	Trims         []string
	Transmissions []string
	Gears         []string
	Extras        []string
}

type storage map[string]common.VDSControl

var store storage

func init() {
	// ensure names are also list in manufacturer.go (getVDSName)
	store = make(storage)
	store["alfa"] = AlfaRomeo{}
	store["audi"] = Audi{}
	store["bmw"] = BMW{}
	store["ferrari"] = Ferrari{}
	store["ford"] = Ford{}
	store["gm"] = GeneralMotors{}
	store["honda"] = Honda{}
	store["hyundai"] = Hyundai{}
	store["isuzu"] = Isuzu{}
	store["kia"] = KIA{}
	store["lamborghini"] = Lamborghini{}
	store["landrover"] = LandRover{}
	store["maserati"] = Maserati{}
	store["mercedes"] = MercedesBenz{}
	store["mitsubishi"] = Mitsubishi{}
	store["seat"] = SEAT{}
	store["skoda"] = Skoda{}
	store["subaru"] = Subaru{}
	store["toyota"] = Toyota{}
	store["volkswagen"] = Volkswagen{}
	store["volvo"] = Volvo{}
}

func GetVDSForBrand(vdsName string) common.VDSControl {
	var result common.VDSControl

	item, ok := store[vdsName]

	if ok {
		result = item
	}

	return result
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func getMany(val string) []string {
	return strings.Split(val, ", ")
}

func fillDirty(dirt *Dirty, data Data) {
	dirt.BodyStyles = appendNoNil(dirt.BodyStyles, getMany(data.BodyStyle)...)
	dirt.Doors = appendNoNil(dirt.Doors, getMany(data.Doors)...)
	dirt.EngineModels = appendNoNil(dirt.EngineModels, data.EngineModel)
	dirt.EngineSizes = appendNoNil(dirt.EngineSizes, getMany(data.EngineSize)...)
	dirt.Extras = appendNoNil(dirt.Extras, getMany(data.Extra)...)
	dirt.Gears = appendNoNil(dirt.Gears, data.Gears)
	dirt.Models = appendNoNil(dirt.Models, data.Model)
	dirt.Transmissions = appendNoNil(dirt.Transmissions, data.Transmission)
	dirt.Trims = appendNoNil(dirt.Trims, getMany(data.Trim)...)
}

func compileDirty(dirt Dirty) common.VDS {
	var result common.VDS

	fmt.Println(dirt)
	result.BodyStyle = getBestGuess(dirt.BodyStyles)
	result.Doors = getBestGuess(dirt.Doors)
	result.EngineModel = getBestGuess(dirt.EngineModels)
	result.EngineSize = getBestGuess(dirt.EngineSizes)
	result.Extras = dirt.Extras
	result.Gears = getBestGuess(dirt.Gears)
	result.Model = getBestGuess(dirt.Models)
	result.Transmission = getBestGuess(dirt.Transmissions)
	result.Trim = getBestGuess(dirt.Trims)

	return result
}

func getBestGuess(items []string) string {
	var result string
	itemLen := len(items)

	if itemLen > 1 {
		mf := 1
		m := 0
		for k, v := range items {
			for i := k; i < len(items); i++ {
				if v == items[i] {
					m++
				}

				if mf < m {
					mf = m
					result = v
				}
			}

			m = 0
		}
	} else if itemLen == 1 {
		result = items[0]
	}

	return result
}

func appendNoNil(arr []string, obj ...string) []string {
	for _, v := range obj {
		if v != "" {
			arr = append(arr, v)
		}
	}

	return arr
}
