package brands

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type SEAT struct {
}

func (b SEAT) GetVDS(sections common.VINSections) common.VDS {
	var result common.VDS

	return result
}

/*
func groupsSEAT() {
	groupv := NewWMIGroup("V")
	groupv.Add("SS", "SEAT", PassengerCar, SEAT{})
}*/
