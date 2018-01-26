package descriptors

type SEAT struct {
}

func (d SEAT) GetData(vinNo string) string {
	return ""
}

func groupsSEAT() {
	groupv := NewWMIGroup("V")
	groupv.Add("SS", "SEAT", PassengerCar, SEAT{})
}
