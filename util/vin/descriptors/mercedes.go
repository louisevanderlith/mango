package descriptors

type MercedesBenz struct {
}

func (d MercedesBenz) GetData(vin string) interface{} {
	return 0
}

func groupsMercedes() {
	descrip := MercedesBenz{}

	groupw := NewWMIGroup("W")
	groupw.Add("DB", "Daimler-Benz/Chrysler", PassengerCar, descrip)
	groupw.Add("DC", "Daimler-Chrysler", PassengerCar, descrip)
	groupw.Add("DD", "Daimler AG", PassengerCar, descrip)
	groupw.Add("DF", "Daimler-Benz", Truck, descrip)
}
