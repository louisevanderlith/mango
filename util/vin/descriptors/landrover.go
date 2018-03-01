package descriptors

type LandRover struct {
}

func (d LandRover) GetData(vinNo string) string {
	return ""
}

func groupsLandRover() {
	const landie = "Land Rover"
	descrip := LandRover{}

	groupa := NewWMIGroup("A")
	groupa.Add("AD", landie, MPV, descrip)

	groups := NewWMIGroup("S")
	groups.Add("AL", landie, MPV, descrip)

	groupw := NewWMIGroup("W")
	groupw.Add("IL", landie, MPV, descrip)

	group9 := NewWMIGroup("9")
	group9.Add("3R", landie, MPV, descrip)
}
