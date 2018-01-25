package descriptors

type Mitsubishi struct {
}

func (d Mitsubishi) GetData(vin string) interface{} {
	return 0
}

func groupsMitsubishi() {
	const mits := "Mitsubishi"
	descrip := Mitsubishi{}

	groupj := NewWMIGroup("J")
	groupj.Add("A3", mits, PassengerCar, descrip)
	groupj.Add("A4", mits, MPV, descrip)
	groupj.Add("A7", mits, Truck, descrip)
	groupj.Add("U3", mits, PassengerCar, descrip)
	groupj.Add("W6", mits, IncompleteCar, descrip)
}
