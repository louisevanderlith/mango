package brands

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type Lamborghini struct {
}

func (b Lamborghini) GetVDS(sections common.VINSections) common.VDS {
	var result common.VDS

	return result
}

/*
func groupsLamborghini() {
	const lambo = "Lamborghini"
	descrip := Lamborghini{}

	groupz := NewWMIGroup("Z")
	groupz.Add("HW", lambo, PassengerCar, descrip)
	groupz.Add("A9", lambo, PassengerCar, descrip)
}*/
