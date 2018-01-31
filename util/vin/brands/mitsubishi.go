package brands

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type Mitsubishi struct {
}

func (b Mitsubishi) GetVDS(sections common.VINSections) common.VDS {
	var result common.VDS

	return result
}

/*
func groupsMitsubishi() {
	const mits = "Mitsubishi"
	descrip := Mitsubishi{}

	groupj := NewWMIGroup("J")
	groupj.Add("A3", mits, PassengerCar, descrip)
	groupj.Add("A4", mits, MPV, descrip)
	groupj.Add("A7", mits, Truck, descrip)
	groupj.Add("U3", mits, PassengerCar, descrip)
	groupj.Add("W6", mits, IncompleteCar, descrip)
}*/
