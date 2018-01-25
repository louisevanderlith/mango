package descriptors

type Audi struct {
}

func (d Audi) GetData(vin string) interface{} {
	return 0
}

func groupsAudi(){
	const audi := "Audi"
	descrip := Audi{}

	groupl := NewWMIGroup("L")
	group1.Add("FV", audi, PassengerCar, descrip)

	groupt := NewWMIGroup("T")
	groupt.Add("RU", audi, PassengerCar, descrip)

	groupw := NewWMIGroup("W")
	groupw.Add("A1", audi, PassengerCar, descrip)
	groupw.Add("AU", audi, PassengerCar, descrip)
	groupw.Add("UA", audi, PassengerCar, descrip)
}
