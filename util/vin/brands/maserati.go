package brands

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type Maserati struct {
}

func (b Maserati) GetVDS(sections common.VINSections) common.VDS {
	var result common.VDS

	return result
}

/*
func groupsMaserati() {
	groupz := NewWMIGroup("Z")
	groupz.Add("AM", "Maserati", PassengerCar, Maserati{})
}*/
