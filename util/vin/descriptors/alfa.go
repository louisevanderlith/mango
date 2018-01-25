package descriptors

type AlfaRomeo struct {
}

func (d AlfaRomeo) GetData(vin string) interface{} {
	return 0
}

func groupsAlfa() {
	groupz := NewWMIGroup("Z")
	groupz.Add("AR", "Alfa Romeo", PassengerCar, AlfaRomeo{})
}
