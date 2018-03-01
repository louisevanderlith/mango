package descriptors

type Skoda struct {
}

func (d Skoda) GetData(vinNo string) string {
	return ""
}

func groupsSkoda() {
	groupt := NewWMIGroup("T")
	groupt.Add("MB", "Skoda", PassengerCar, Skoda{})
}
