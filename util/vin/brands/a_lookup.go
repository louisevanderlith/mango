package brands

import "github.com/louisevanderlith/mango/util/vin/common"

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
