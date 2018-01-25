package descriptors

type BMW struct {
}

func (d BMW) GetData(vin string) interface{} {
	return 0
}

func groupsBMW(){
	const bmw := "BMW"
	const mini := "Mini"
	const rolls := "Rolls Royce"
	descrip := BMW{}

	groupm := NewWMIGroup("M")
	groupm.Add("MF", bmw, PassengerCar, descrip)

	groupn := NewWMIGroup("N")
	groupn.Add("C0", bmw, PassengerCar, descrip)

	groups := NewWMIGroup("S")
	groups.Add("A9", rolls, PassengerCar, descrip)
	groups.Add("CA", rolls, PassengerCar, descrip)

	groupw := NewWMIGroup("W")
	groupw.Add("BA", bmw, PassengerCar, descrip)
	groupw.Add("BS", bmw, PassengerCar, descrip)
	groupw.Add("B1", bmw, Motorcycle, descrip)
	groupw.Add("MW", mini, PassengerCar, descrip)
	groupw.Add("BX", bmw, PassengerCar, descrip)

	groupx := NewWMIGroup("X")
	groupx.Add("4X", bmw, PassengerCar, descrip)

	group3 := NewWMIGroup("3")
	group3.Add("AV", bmw, PassengerCar, descrip)

	group4 := NewWMIGroup("4")
	group4.Add("US", bmw, PassengerCar, descrip)

	group5 := NewWMIGroup("5")
	group5.Add("UM", bmw, PassengerCar, descrip)
	group5.Add("UX", bmw, PassengerCar, descrip)
}
