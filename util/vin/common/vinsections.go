package common

import "errors"

type VINSections struct {
	WMICode           string
	FeatureCode       string
	CheckDigit        string
	YearCode          string
	AssemblyPlantCode string
	SequenceCode      string
	FullVIN           string
}

func LoadVINSections(vinNo string) (result VINSections, err error) {
	if len(vinNo) == 17 {
		result.FullVIN = vinNo
		result.WMICode = vinNo[:3]
		result.FeatureCode = vinNo[3:8]
		result.CheckDigit = vinNo[8:9]
		result.YearCode = vinNo[9:10]
		result.AssemblyPlantCode = vinNo[10:11]
		result.SequenceCode = vinNo[11:]
	} else {
		err = errors.New("VIN is not 17 characters")
	}

	return result, err
}
