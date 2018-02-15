package vin

import (
	"errors"

	"github.com/louisevanderlith/mango/util/vin/brands"
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
		result.VIS = loadVIS(sections)

		if result.VIS.ValidVIN {
			result.VDS = loadVDS(sections, result.WMI.Manufacturer.VDSName, result.VIS.Year)
		} else {
			err = errors.New("vin is not valid")
		}
	}

	return result, err
}

func loadWMI(sections common.VINSections) common.WMI {
	var result common.WMI

	result.Region = common.GetRegion(sections.WMICode)
	result.Manufacturer = common.GetManufacturer(sections.WMICode)

	return result
}

func loadVDS(sections common.VINSections, vdsName string, year int) common.VDS {
	var result common.VDS

	brandVDS := brands.GetVDSForBrand(vdsName)

	result = brandVDS.GetPassengerCar(sections, year)

	return result
}

func loadVIS(sections common.VINSections) common.VIS {
	var result common.VIS

	result.ValidVIN = common.IsValid(sections.FullVIN, sections.VISCode.CheckDigit)
	result.Year = common.GetYear(sections.VISCode.YearCode, sections.VDSCode.Char7)
	result.SequenceNo = sections.VISCode.SequenceCode

	return result
}
