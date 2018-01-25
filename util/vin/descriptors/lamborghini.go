package descriptors

type Lamborghini struct {
}

func (d Lamborghini) GetData(vin string) interface{} {
	return 0
}

func groupsLamborghini() {
	const lambo := "Lamborghini"
	descrip := Lamborghini{}

	groupz := NewWMIGroup("Z")
	groupz.Add("HW", lambo, PassengerCar, descrip)
	groupz.Add("A9", lambo, PassengerCar, descrip)
}
