package descriptors

type Subaru struct {
}

func (d Subaru) GetData(vin string) interface{} {
	return 0
}

func groupsSubaru(){
	const fuji := "Fuji Heavy Industries"
	const subaru := "Subaru"
	descrip := Subaru{}

	groupj := NewWMIGroup("J")
	groupj.Add("F1", fuji, PassengerCar, descrip)
	groupj.Add("F2", fuji, MPV, descrip)

	group4 := NewWMIGroup("4")
	group4.Add("S3", subaru, PassengerCar, descrip)
	group4.Add("S4", subaru, MPV, descrip)
}
