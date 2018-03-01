package descriptors

type Audi struct {
}

func (d Audi) GetData(vinNo string) string {
	return ""
}

func groupsAudi() {
	const audi = "Audi"
	descrip := Audi{}

	group1 := NewWMIGroup("L")
	group1.Add("FV", audi, PassengerCar, descrip)

	groupt := NewWMIGroup("T")
	groupt.Add("RU", audi, PassengerCar, descrip)

	groupw := NewWMIGroup("W")
	groupw.Add("A1", audi, PassengerCar, descrip)
	groupw.Add("AU", audi, PassengerCar, descrip)
	groupw.Add("UA", audi, PassengerCar, descrip)
}
