package vin

import (
	"fmt"

	"github.com/louisevanderlith/mango/util/vin/common"
	"github.com/louisevanderlith/mango/util/vin/section"
)

type Info struct {
	WMI section.WMI
	VDS section.VDS
	VIS section.VIS
}

// GetInfo is the main entry-point for reading VIN information
func GetInfo(vinNo string) (result Info, err error) {
	sections, err := common.LoadVINSections(vinNo)

	if err == nil {
		fmt.Print(sections)
		result.WMI = section.LoadWMI(sections)
		result.VDS = section.LoadVDS(sections)
		result.VIS = section.LoadVIS(sections)
	}

	return result, err
}
