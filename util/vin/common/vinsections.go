package common

import "errors"

type VINSections struct {
	WMICode WMICode
	VDSCode VDSCode
	VISCode VISCode
	FullVIN string
}

type WMICode struct {
	FullWMI          string
	ContinentCode    string
	RegionCode       string
	ManufacturerCode string
}

type VDSCode struct {
	FullVDS string
	Char4   string
	Char5   string
	Char6   string
	Char7   string
	Char8   string
}

func (v VDSCode) GetTypeCode(characters ...string) string {
	var result string

	codes := make(map[string]string)
	codes["4"] = v.Char4
	codes["5"] = v.Char5
	codes["6"] = v.Char6
	codes["7"] = v.Char7
	codes["8"] = v.Char8

	for _, v := range characters {
		if char, ok := codes[v]; ok {
			result += char
		}
	}

	return result
}

type VISCode struct {
	FullVIS           string
	CheckDigit        string
	YearCode          string
	AssemblyPlantCode string
	SequenceCode      string
}

func LoadVINSections(vinNo string) (result VINSections, err error) {
	if len(vinNo) == 17 {
		result.FullVIN = vinNo
		result.WMICode = WMICode{
			FullWMI:          vinNo[:3],
			ContinentCode:    vinNo[:1],
			RegionCode:       vinNo[1:2],
			ManufacturerCode: vinNo[1:3],
		}
		result.VDSCode = VDSCode{
			FullVDS: vinNo[3:8],
			Char4:   vinNo[3:4],
			Char5:   vinNo[4:5],
			Char6:   vinNo[5:6],
			Char7:   vinNo[6:7],
			Char8:   vinNo[7:8],
		}

		result.VISCode = VISCode{
			FullVIS:           vinNo[8:],
			CheckDigit:        vinNo[8:9],
			YearCode:          vinNo[9:10],
			AssemblyPlantCode: vinNo[10:11],
			SequenceCode:      vinNo[11:],
		}
	} else {
		err = errors.New("VIN is not 17 characters")
	}

	return result, err
}
