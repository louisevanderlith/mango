package brands

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type Skoda struct {
}

func (b Skoda) GetVDS(sections common.VINSections) common.VDS {
	var result common.VDS

	return result
}

/*
func groupsSkoda() {
	groupt := NewWMIGroup("T")
	groupt.Add("MB", "Skoda", PassengerCar, Skoda{})
}*/
