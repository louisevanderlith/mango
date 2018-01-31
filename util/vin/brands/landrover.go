package brands

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type LandRover struct {
}

func (b LandRover) GetVDS(sections common.VINSections) common.VDS {
	var result common.VDS

	return result
}

/*
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
*/
