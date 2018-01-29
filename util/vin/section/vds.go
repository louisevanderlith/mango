package section

import (
	"fmt"

	"github.com/louisevanderlith/mango/util/vin/common"
	"github.com/louisevanderlith/mango/util/vin/descriptors"
)

type VDS struct {
	BodyStyle string
	Engine    string
	Restraint string
}

func LoadVDS(sections common.VINSections) VDS {
	var result VDS

	wmiDescrip := descriptors.GetWMIDescriptor(sections.WMICode)

	if wmiDescrip != (descriptors.WMIDescriptor{}) {
		data := wmiDescrip.Descriptor.GetData(sections.FeatureCode)
		fmt.Println(data)
	}

	return result
}
