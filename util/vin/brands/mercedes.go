package brands

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type MercedesBenz struct {
}

func (b MercedesBenz) GetVDS(sections common.VINSections) common.VDS {
	var result common.VDS

	return result
}

/*
func groupsMercedes() {
	descrip := MercedesBenz{}

	groupw := NewWMIGroup("W")
	groupw.Add("DB", "Daimler-Benz/Chrysler", PassengerCar, descrip)
	groupw.Add("DC", "Daimler-Chrysler", PassengerCar, descrip)
	groupw.Add("DD", "Daimler AG", PassengerCar, descrip)
	groupw.Add("DF", "Daimler-Benz", Truck, descrip)
}*/
