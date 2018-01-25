package descriptors

type Isuzu struct {
}

func (d Isuzu) GetData(vin string) interface{} {
	return 0
}

func groupsIsuzu() {
	const isuzu := "Isuzu"
	descrip := Isuzu{}

	groupj := NewWMIGroup("J")
	groupj.Add("AA", isuzu, MPV, descrip)
	groupj.Add("AB", isuzu, MPV, descrip)
	groupj.Add("AC", isuzu, MPV, descrip)
	groupj.Add("AL", isuzu, IncompleteCar, descrip)
	groupj.Add("81", isuzu, MPV, descrip)
	groupj.Add("82", isuzu, MPV, descrip)
}
