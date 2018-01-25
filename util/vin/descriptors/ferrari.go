package descriptors

type Ferrari struct {
}

func (d Ferrari) GetData(vin string) interface{} {
	return 0
}

func groupsFerrari(){
	const ferrari := "Ferrari"
	descrip := Ferrari{}

	groupz := NewWMIGroup("Z")
	groupz.Add("DF", ferrari, PassengerCar, descrip)
	groupz.Add("FF",ferrari, PassengerCar, descrip)
}
