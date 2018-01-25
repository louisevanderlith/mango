package descriptors

type Ford struct {

}

func (d Ford) GetData(vin string) interface{
	return 0
}

func groupsFord(){
	const ford := "Ford"
	const lincoln := "Lincoln"
	const mercury := "Mercury"
	const mazda := "Mazda"
	descrip := Ford{}

	groupa := NewWMIGroup("A")
	groupa.Add("FA", ford, NotSpecified, descrip)

	groupm := NewWMIGroup("M")
	groupm.Add("PB", ford, NotSpecified, descrip)
	groupm.Add("NB", ford, NotSpecified, descrip)

	groupn := NewWMIGroup("N")
	groupn.Add("M0", ford, NotSpecified, descrip)
	
	groupp := NewWMIGroup("P")
	groupp.Add("E1", ford, NotSpecified, descrip)

	groups := NewWMIGroup("S")
	groups.Add("FA", ford, NotSpecified, descrip)

	groupt := NewWMIGroup("T")
	groupt.Add("W2", ford, NotSpecified, descrip)

	groupu := NewWMIGroup("U")
	groupu.Add("NI", ford, NotSpecified, descrip)

	groupv := NewWMIGroup("V")
	groupv.Add("S6", ford, NotSpecified, descrip)

	groupw := NewWMIGroup("W")
	groupw.Add("F0", ford, NotSpecified, descrip)
	groupw.Add("F1", ford, NotSpecified, descrip)

	groupx := NewWMIGroup("X")
	groupx.Add("LC", ford, NotSpecified, descrip)

	groupy := NewWMIGroup("Y")
	groupy.Add("4F", ford, NotSpecified, descrip)

	group1 := NewWMIGroup("1")
	group1.Add("F1", ford, MPV, descrip)
	group1.Add("F6", ford, BasicChassis, descrip)
	group1.Add("F7", ford, PassengerCar, descrip)
	group1.Add("FA", ford, PassengerCar, descrip)
	group1.Add("FB", ford, Bus, descrip)
	group1.Add("FC", ford, BasicChassis, descrip)
	group1.Add("FD", ford, IncompleteCar, descrip)
	group1.Add("FM", ford, MPV, descrip)
	group1.Add("FT", ford, Truck, descrip)
	group1.Add("L1", lincoln, PassengerCar, descrip)
	group1.Add("LJ", lincoln, IncompleteCar, descrip)
	group1.Add("LN", lincoln, PassengerCar, descrip)
	group1.Add("ME", mercury, PassengerCar, descrip)
	group1.Add("MH", mercury, IncompleteCar, descrip)
	group1.Add("YV", mazda, PassengerCar, descrip)
	group1.Add("ZV", ford, PassengerCar, descrip)

	group2 := NewWMIGroup("2")
	group2.Add("FA", ford, PassengerCar, descrip)
	group2.Add("FD", ford, IncompleteCar, descrip)
	group2.Add("FM", ford, MPV, descrip)
	group2.Add("FT", ford, Truck, descrip)
	group2.Add("ME", mercury, PassengerCar, descrip)
	group2.Add("MH", mercury, IncompleteCar, descrip)
	group2.Add("MR", mercury, MPV, descrip)

	group3 := NewWMIGroup("3")
	group3.Add("FA", ford, PassengerCar, descrip)
	group3.Add("FD", ford, IncompleteCar, descrip)
	group3.Add("FT", ford, Truck, descrip)
	group3.Add("MA", mercury, PassengerCar, descrip)
	group3.Add("FN", ford, Truck, descrip)
	group3.Add("FR", ford, IncompleteCar, descrip)

	group4 := NewWMIGroup("4")
	group4.Add("F2", mazda, MPV, descrip)
	group4.Add("F4", mazda, Truck, descrip)
	group4.Add("M2", mercury, MPV, descrip)
	
	group5 := NewWMIGroup("5")
	group5.Add("L1", lincoln, MPV, descrip)
	group5.Add("LT", lincoln, Truck, descrip)

	group6 := NewWMIGroup("6")
	group6.Add("FP", ford, NotSpecified, descrip)

	group8 := NewWMIGroup("8")
	group8.Add("AF", ford, NotSpecified, descrip)
	group8.Add("XD", ford, NotSpecified, descrip)

	group9 := NewWMIGroup("9")
	group9.Add("BF", ford, NotSpecified, descrip)
}

