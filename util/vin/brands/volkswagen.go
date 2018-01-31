package brands

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type Volkswagen struct {
}

func (b Volkswagen) GetVDS(sections common.VINSections) common.VDS {
	var result common.VDS

	return result
}

/*
func groupsVW() {
	const vw = "Volkswagen"
	descrip := Volkswagen{}

	groupa := NewWMIGroup("A")
	groupa.Add("AV", vw, PassengerCar, descrip)

	groupl := NewWMIGroup("L")
	groupl.Add("SV", vw, PassengerCar, descrip)

	groupw := NewWMIGroup("W")
	groupw.Add("VW", vw, PassengerCar, descrip)
	groupw.Add("V1", vw, PassengerCar, descrip)
	groupw.Add("V2", vw, PassengerCar, descrip)
	groupw.Add("V3", vw, PassengerCar, descrip)
	groupw.Add("VG", vw, PassengerCar, descrip)

	groupx := NewWMIGroup("X")
	groupx.Add("W8", vw, PassengerCar, descrip)

	group1 := NewWMIGroup("1")
	group1.Add("VW", vw, PassengerCar, descrip)

	group2 := NewWMIGroup("2")
	group2.Add("V4", vw, PassengerCar, descrip)
	group2.Add("V8", vw, PassengerCar, descrip)

	group3 := NewWMIGroup("3")
	group3.Add("VW", vw, PassengerCar, descrip)

	group8 := NewWMIGroup("8")
	group8.Add("AW", vw, PassengerCar, descrip)

	group9 := NewWMIGroup("9")
	group9.Add("3U", vw, PassengerCar, descrip)
	group9.Add("BW", vw, PassengerCar, descrip)
}
*/
