package descriptors

type SEAT struct {
}

func (d SEAT) GetData(vin string) interface{} {
	return 0
}

func groupsSEAT() {
	groupv := NewWMIGroup("V")
	groupv.Add("SS", "SEAT", PassengerCar, SEAT{})
}
