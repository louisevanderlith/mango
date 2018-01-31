package brands

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type KIA struct {
}

func (b KIA) GetVDS(sections common.VINSections) common.VDS {
	var result common.VDS

	return result
}

/*
func groupsKIA() {
	groupm := NewWMIGroup("M")
	groupm.Add("S0", "KIA", NotSpecified, KIA{})
}*/
