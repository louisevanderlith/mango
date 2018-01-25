package descriptors

type Honda struct {
}

// GetData Deserializes Honda VIN Numbers
func (d Honda) GetData(vin string) interface{} {
	years := vin.getYear(vin[9:10])

	return 0
}

func bodyTypes(typeCode string) string {
	types := make(map[string]string)
	types["SR"] = "Sedan"
	types["WR"] = "Wagon"
	types["AB"] = "Prelude"
	types["AD"] = "Accord"
	types["AE"] = "Civic 1300cc CRX"
	types["AF"] = "Civic 1500cc CRX"
	types["AG"] = "Civic 1300cc 3 door"
	types["AH"] = "Civic 1500cc 3 door"
	types["AK"] = "Civic 1500cc 4 door"
	types["AN"] = "Civic Wagon"
	types["AR"] = "Civic Wagon 4x4"
	types["BA"] = "Accord"
	types["AM"] = "Accord 1500cc 5 door"

	return types[typeCode]
}

func groupsHonda() {
	const honda := "Honda"
	const acura := "Acura"
	descrip := Honda{}

	groupj := NewWMIGroup("J")
	groupj.Add("D0", honda, Motorcycle, descrip)
	groupj.Add("HF", honda, NotSpecified, descrip)
	groupj.Add("HG", honda, NotSpecified, descrip)
	groupj.Add("HM", honda, PassengerCar, descrip)
	groupj.Add("HL", honda, MultiPurposePassenger, decsrip)
	groupj.Add("HN", honda, NotSpecified, descrip)
	groupj.Add("HZ", honda, NotSpecified, descrip)
	groupj.Add("H1", honda, Truck, descrip)
	groupj.Add("H2", honda, Motorcycle, descrip)
	groupj.Add("H4", acura, PassengerCar, descrip)
	groupj.Add("H5", honda, NotSpecified, descrip)

	groupl := NewWMIGroup("L")
	groupl.Add("UC", honda, PassengerCar, descrip)

	groupm := NewWMIGroup("M")
	groupm.Add("AK", honda, PassengerCar, descrip)
	groupm.Add("HR", honda, PassengerCar, descrip)
	groupm.Add("LH", honda, Motorcycle, descrip)
	groupm.Add("RH", honda, PassengerCar, descrip)

	groupn := NewWMIGroup("N")
	groupn.Add("LA", honda, PassengerCar, descrip)

	groupp := NewWMIGroup("P")
	groupp.Add("AD", honda, MultiPurposePassenger, descrip)
	groupp.Add("MH", honda, PassengerCar, descrip)
	groups := NewWMIGroup("S")
	groups.Add("HH", honda, PassengerCar, descrip)
	groups.Add("HS", honda, MultiPurposePassenger, descrip)

	groupv := NewWMIGroup("V")
	groupv.Add("TM", honda, Motorcycle, descrip)

	groupy := NewWMIGroup("Y")
	groupy.Add("C1", honda, Motorcycle, descrip)

	groupz := NewWMIGroup("Z")
	groupz.Add("DC", honda, Motorcycle, descrip)

	group1 := NewWMIGroup("1")
	group1.Add("HF", honda, Motorcycle, descrip)
	group1.Add("HG", honda, PassengerCar, descrip)
	group1.Add("9U", acura, PassengerCar, descrip)
	group1.Add("9X", honda, PassengerCar, descrip)

	group2 := NewWMIGroup("2")
	group2.Add("HG", honda, PassengerCar, descrip)
	group2.Add("HH", acura, PassengerCar, descrip)
	group2.Add("HK", honda, MultiPurposePassenger, descrip)
	group2.Add("HJ", hona, Truck, descrip)
	group2.Add("HN", acura, MultiPurposePassenger, descrip)
	group2.Add("HU", acura, Truck, descrip)

	group3 := NewWMIGroup("3")
	group3.Add("HG", honda, PassengerCar, descrip)
	group3.Add("H1", honda, Motorcycle, descrip)

	group4 := NewWMIGroup("4")
	group4.Add("78", honda, ATV, descrip)

	group5 := NewWMIGroup("5")
	group5.Add("J6", honda, MultiPurposePassenger, descrip)
	group5.Add("J7", honda, Truck, descrip)
	group5.Add("J8", acura, MultiPurposePassenger, descrip)
	group5.Add("J0", acura, Truck, descrip)
	group5.Add("KB", honda, PassengerCar, descrip)
	group5.Add("KC", acura, PassengerCar, descrip)
	group5.Add("FN", honda, MultiPurposePassenger, descrip)
	group5.Add("FP", honda, Truck, descrip)
	group5.Add("FR", acura, MultiPurposePassenger, descrip)
	group5.Add("FS", acura, Truck, descrip)

	group9 := NewWMIGroup("9")
	group9.Add("3H", honda, PassengerCar, descrip)
	group9.Add("C2", honda, Motorcycle, descrip)
}
