package descriptors

type Maserati struct {
}

func (d Maserati) GetData(vinNo string) string {
	return ""
}

func groupsMaserati() {
	groupz := NewWMIGroup("Z")
	groupz.Add("AM", "Maserati", PassengerCar, Maserati{})
}
