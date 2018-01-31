package vin

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type Info struct {
	WMI common.WMI
	VDS common.VDS
	VIS common.VIS
}

// GetInfo is the main entry-point for reading VIN information
func GetInfo(vinNo string) (result Info, err error) {
	sections, err := common.LoadVINSections(vinNo)

	if err == nil {
		result.WMI = loadWMI(sections)
		result.VDS = loadVDS(sections)
		result.VIS = loadVIS(sections)
	}

	return result, err
}

func loadWMI(sections common.VINSections) common.WMI {
	var result common.WMI

	result.Region = common.GetRegion(sections.ContinentCode, sections.RegionCode)
	result.Manufacturer = common.GetManufacturer(sections.ContinentCode, sections.ManufacturerCode)

	return result
}

func loadVDS(sections common.VINSections) common.VDS {
	var result common.VDS
	// TODO
	return result
}

func loadVIS(sections common.VINSections) common.VIS {
	var result common.VIS

	result.ValidVIN = common.IsValid(sections.FullVIN, sections.CheckDigit)
	result.Year = common.GetBGYear(sections.YearCode)
	result.AssemblyPlant = ""
	result.SequenceNo = sections.SequenceCode

	return result
}
