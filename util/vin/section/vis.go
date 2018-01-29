package section

import "github.com/louisevanderlith/mango/util/vin/common"

type VIS struct {
	ValidVIN      bool
	Year          int
	AssemblyPlant string
	SequenceNo    string
}

func LoadVIS(sections common.VINSections) VIS {
	var result VIS

	result.ValidVIN = common.IsValid(sections.FullVIN, sections.CheckDigit)
	result.Year = common.GetBGYear(sections.YearCode)
	result.AssemblyPlant = ""
	result.SequenceNo = sections.SequenceCode

	return result
}
