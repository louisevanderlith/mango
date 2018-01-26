package descriptors

type Ferrari struct {
}

func (d Ferrari) GetData(vinNo string) string {
	return ""
}

func groupsFerrari() {
	const ferrari = "Ferrari"
	descrip := Ferrari{}

	groupz := NewWMIGroup("Z")
	groupz.Add("DF", ferrari, PassengerCar, descrip)
	groupz.Add("FF", ferrari, PassengerCar, descrip)
}
