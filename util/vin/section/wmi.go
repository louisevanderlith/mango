package section

import "github.com/louisevanderlith/mango/util/vin/common"

// WMI is the World Manufacturer Indentifier
type WMI struct {
	Region       common.Region
	Manufacturer common.Manufacturer
}

func LoadWMI(sections common.VINSections) WMI {
	var result WMI

	if len(sections.WMICode) == 3 {
		regionCode := sections.WMICode[:2]
		manCode := sections.WMICode[2:3]

		result.Region = common.GetRegion(regionCode)
		result.Manufacturer = common.GetManufacturer(manCode)
	}

	return result
}
