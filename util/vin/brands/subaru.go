package brands

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type Subaru struct {
}

func (b Subaru) GetVDS(sections common.VINSections) common.VDS {
	var result common.VDS

	return result
}

/*
func groupsSubaru() {
	const fuji = "Fuji Heavy Industries"
	const subaru = "Subaru"
	descrip := Subaru{}

	groupj := NewWMIGroup("J")
	groupj.Add("F1", fuji, PassengerCar, descrip)
	groupj.Add("F2", fuji, MPV, descrip)

	group4 := NewWMIGroup("4")
	group4.Add("S3", subaru, PassengerCar, descrip)
	group4.Add("S4", subaru, MPV, descrip)
}*/
