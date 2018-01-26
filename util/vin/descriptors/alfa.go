package descriptors

type AlfaRomeo struct {
}

func (d AlfaRomeo) GetData(vinNo string) string {
	return ""
}

func groupsAlfa() {
	groupz := NewWMIGroup("Z")
	groupz.Add("AR", "Alfa Romeo", PassengerCar, AlfaRomeo{})
}
