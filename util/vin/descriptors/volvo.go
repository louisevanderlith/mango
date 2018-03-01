package descriptors

type Volvo struct {
}

func (d Volvo) GetData(vinNo string) string {
	return ""
}

func groupsVolvo() {
	const volvo = "Volvo"
	descrip := Volvo{}

	groupx := NewWMIGroup("X")
	groupx.Add("LB", volvo, PassengerCar, descrip)

	groupy := NewWMIGroup("Y")
	groupy.Add("V1", volvo, PassengerCar, descrip)
	groupy.Add("V2", volvo, Truck, descrip)
	groupy.Add("V3", volvo, Bus, descrip)
	groupy.Add("V4", volvo, MPV, descrip)

	group4 := NewWMIGroup("4")
	group4.Add("V1", volvo, Truck, descrip)
	group4.Add("V2", volvo, NotSpecified, descrip)
	group4.Add("V3", volvo, NotSpecified, descrip)
	group4.Add("V4", volvo, Truck, descrip)
	group4.Add("V5", volvo, Truck, descrip)
	group4.Add("V6", volvo, NotSpecified, descrip)
	group4.Add("VL", volvo, NotSpecified, descrip)
	group4.Add("VM", volvo, NotSpecified, descrip)
	group4.Add("VZ", volvo, NotSpecified, descrip)

	groupm := NewWMIGroup("M")
	groupm.Add("HA", volvo, NotSpecified, descrip)
}
