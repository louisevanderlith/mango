package brands

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type Ford struct {
}

func (b Ford) GetVDS(sections common.VINSections) common.VDS {
	var result common.VDS

	return result
}

/*
func groupsFord() {
	const ford = "Ford"
	const lincoln = "Lincoln"
	const mercury = "Mercury"
	const mazda = "Mazda"
	descrip := Ford{}

	groupa := section.NewWMIGroup("A")
	groupa.Add("FA", ford, section.NotSpecified, descrip)

	groupm := section.NewWMIGroup("M")
	groupm.Add("PB", ford, section.NotSpecified, descrip)
	groupm.Add("NB", ford, section.NotSpecified, descrip)

	groupn := section.NewWMIGroup("N")
	groupn.Add("M0", ford, section.NotSpecified, descrip)

	groupp := section.NewWMIGroup("P")
	groupp.Add("E1", ford, section.NotSpecified, descrip)

	groups := section.NewWMIGroup("S")
	groups.Add("FA", ford, section.NotSpecified, descrip)

	groupt := section.NewWMIGroup("T")
	groupt.Add("W2", ford, section.NotSpecified, descrip)

	groupu := section.NewWMIGroup("U")
	groupu.Add("NI", ford, section.NotSpecified, descrip)

	groupv := section.NewWMIGroup("V")
	groupv.Add("S6", ford, section.NotSpecified, descrip)

	groupw := section.NewWMIGroup("W")
	groupw.Add("F0", ford, section.NotSpecified, descrip)
	groupw.Add("F1", ford, section.NotSpecified, descrip)

	groupx := section.NewWMIGroup("X")
	groupx.Add("LC", ford, section.NotSpecified, descrip)

	groupy := section.NewWMIGroup("Y")
	groupy.Add("4F", ford, section.NotSpecified, descrip)

	group1 := section.NewWMIGroup("1")
	group1.Add("F1", ford, section.MPV, descrip)
	group1.Add("F6", ford, section.BasicChassis, descrip)
	group1.Add("F7", ford, section.PassengerCar, descrip)
	group1.Add("FA", ford, section.PassengerCar, descrip)
	group1.Add("FB", ford, section.Bus, descrip)
	group1.Add("FC", ford, section.BasicChassis, descrip)
	group1.Add("FD", ford, section.IncompleteCar, descrip)
	group1.Add("FM", ford, section.MPV, descrip)
	group1.Add("FT", ford, section.Truck, descrip)
	group1.Add("L1", lincoln, section.PassengerCar, descrip)
	group1.Add("LJ", lincoln, section.IncompleteCar, descrip)
	group1.Add("LN", lincoln, section.PassengerCar, descrip)
	group1.Add("ME", mercury, section.PassengerCar, descrip)
	group1.Add("MH", mercury, section.IncompleteCar, descrip)
	group1.Add("YV", mazda, section.PassengerCar, descrip)
	group1.Add("ZV", ford, section.PassengerCar, descrip)

	group2 := section.NewWMIGroup("2")
	group2.Add("FA", ford, section.PassengerCar, descrip)
	group2.Add("FD", ford, section.IncompleteCar, descrip)
	group2.Add("FM", ford, section.MPV, descrip)
	group2.Add("FT", ford, section.Truck, descrip)
	group2.Add("ME", mercury, section.PassengerCar, descrip)
	group2.Add("MH", mercury, section.IncompleteCar, descrip)
	group2.Add("MR", mercury, section.MPV, descrip)

	group3 := section.NewWMIGroup("3")
	group3.Add("FA", ford, section.PassengerCar, descrip)
	group3.Add("FD", ford, section.IncompleteCar, descrip)
	group3.Add("FT", ford, section.Truck, descrip)
	group3.Add("MA", mercury, section.PassengerCar, descrip)
	group3.Add("FN", ford, section.Truck, descrip)
	group3.Add("FR", ford, section.IncompleteCar, descrip)

	group4 := section.NewWMIGroup("4")
	group4.Add("F2", mazda, section.MPV, descrip)
	group4.Add("F4", mazda, section.Truck, descrip)
	group4.Add("M2", mercury, section.MPV, descrip)

	group5 := section.NewWMIGroup("5")
	group5.Add("L1", lincoln, section.MPV, descrip)
	group5.Add("LT", lincoln, section.Truck, descrip)

	group6 := section.NewWMIGroup("6")
	group6.Add("FP", ford, section.NotSpecified, descrip)

	group8 := section.NewWMIGroup("8")
	group8.Add("AF", ford, section.NotSpecified, descrip)
	group8.Add("XD", ford, section.NotSpecified, descrip)

	group9 := section.NewWMIGroup("9")
	group9.Add("BF", ford, section.NotSpecified, descrip)
}
*/
