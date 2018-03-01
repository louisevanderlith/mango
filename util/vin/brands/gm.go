package brands

import "github.com/louisevanderlith/mango/util/vin/common"

type GeneralMotors struct {
	common.VDS
}

/*
func groupsGM() {
	const chev = "Chevrolet"
	const daewoo = "GM Daewoo"
	const pontiac = "Pontiac"
	const olds = "Oldsmobile"
	const buick = "Buick"
	const caddi = "Cadillac"
	const saturn = "Saturn"
	const gmc = "GMC"

	descrip := GeneralMotors{}

	groupk := NewWMIGroup("K")
	groupk.Add("LA", daewoo, PassengerCar, descrip)
	groupk.Add("L1", chev, PassengerCar, descrip)
	groupk.Add("L2", pontiac, PassengerCar, descrip)

	groupw := NewWMIGroup("W")
	groupw.Add("06", caddi, PassengerCar, descrip)
	groupw.Add("08", saturn, PassengerCar, descrip)

	group1 := NewWMIGroup("1")
	group1.Add("GT", gmc, MPV, descrip)
	group1.Add("G1", chev, PassengerCar, descrip)
	group1.Add("G2", pontiac, PassengerCar, descrip)
	group1.Add("G3", olds, PassengerCar, descrip)
	group1.Add("G4", buick, PassengerCar, descrip)
	group1.Add("G5", pontiac, PassengerCar, descrip)
	group1.Add("G6", caddi, PassengerCar, descrip)
	group1.Add("G8", saturn, PassengerCar, descrip)

	group2 := NewWMIGroup("2")
	group2.Add("G1", chev, PassengerCar, descrip)
	group2.Add("G2", pontiac, PassengerCar, descrip)
	group2.Add("G3", olds, PassengerCar, descrip)
	group2.Add("G4", buick, PassengerCar, descrip)
	group2.Add("G6", caddi, PassengerCar, descrip)

	group3 := NewWMIGroup("3")
	group3.Add("G1", chev, PassengerCar, descrip)
	group3.Add("G2", pontiac, PassengerCar, descrip)
	group3.Add("G4", buick, PassengerCar, descrip)

	group5 := NewWMIGroup("5")
	group5.Add("Y2", pontiac, PassengerCar, descrip)

	group6 := NewWMIGroup("6")
	group6.Add("G2", pontiac, PassengerCar, descrip)

	group8 := NewWMIGroup("8")
	group8.Add("AG", chev, PassengerCar, descrip)
	group8.Add("GG", chev, PassengerCar, descrip)
	group8.Add("Z1", chev, PassengerCar, descrip)

	group9 := NewWMIGroup("9")
	group9.Add("BG", chev, PassengerCar, descrip)

}
*/
