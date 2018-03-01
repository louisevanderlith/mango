package descriptors

type Hyundai struct {
}

func (d Hyundai) GetData(vinNo string) string {
	return ""
}

func groupsHyundai() {
	const hyundai = "Hyundai"
	const kia = "Kia"
	descrip := Hyundai{}

	groupa := NewWMIGroup("A")
	groupa.Add("C5", hyundai, PassengerCar, descrip)
	groupa.Add("DD", hyundai, MPV, descrip)

	groupk := NewWMIGroup("K")
	groupk.Add("MF", hyundai, MPV, descrip)
	groupk.Add("MH", hyundai, PassengerCar, descrip)
	groupk.Add("MJ", hyundai, Bus, descrip)
	groupk.Add("MX", hyundai, NotSpecified, descrip)
	groupk.Add("MZ", hyundai, NotSpecified, descrip)
	groupk.Add("M8", hyundai, MPV, descrip)
	groupk.Add("ND", hyundai, MPV, descrip)

	groupl := NewWMIGroup("L")
	groupl.Add("BE", hyundai, NotSpecified, descrip)

	groupm := NewWMIGroup("M")
	groupm.Add("AL", hyundai, NotSpecified, descrip)

	groupn := NewWMIGroup("N")
	groupn.Add("LH", hyundai, PassengerCar, descrip)
	groupn.Add("LJ", hyundai, Bus, descrip)

	groupp := NewWMIGroup("P")
	groupp.Add("EI", hyundai, PassengerCar, descrip)

	groupt := NewWMIGroup("T")
	groupt.Add("MA", hyundai, NotSpecified, descrip)

	groupu := NewWMIGroup("U")
	groupu.Add("6Z", kia, PassengerCar, descrip)

	groupz := NewWMIGroup("Z")
	groupz.Add("94", hyundai, NotSpecified, descrip)
	groupz.Add("K5", hyundai, NotSpecified, descrip)

	group2 := NewWMIGroup("2")
	group2.Add("HM", hyundai, NotSpecified, descrip)

	group3 := NewWMIGroup("3")
	group3.Add("H3", hyundai, NotSpecified, descrip)

	group5 := NewWMIGroup("5")
	group5.Add("NM", hyundai, MPV, descrip)
	group5.Add("NP", hyundai, PassengerCar, descrip)

	group9 := NewWMIGroup("9")
	group9.Add("5P", hyundai, MPV, descrip)
	group9.Add("5*", hyundai, Truck, descrip)
	group9.Add("BH", hyundai, NotSpecified, descrip)
}
