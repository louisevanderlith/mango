package descriptors

type Maserati struct {
}

func (d Maserati) GetData(vin string) interface{} {
	return 0
}

func groupsMaserati() {
	groupz := NewWMIGroup("Z")
	groupz.Add("AM", "Maserati", PassengerCar, Maserati{})
}
