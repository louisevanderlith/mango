package descriptors

type Skoda struct {
}

func (d Skoda) GetData(vin string) interface{} {
	return 0
}

func groupsSkoda() {
	groupt := NewWMIGroup("T")
	groupt.Add("MB", "Skoda", PassengerCar, Skoda{})
}
