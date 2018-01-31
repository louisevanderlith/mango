package brands

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type Isuzu struct {
}

func (b Isuzu) GetVDS(sections common.VINSections) common.VDS {
	var result common.VDS

	return result
}

/*
func groupsIsuzu() {
	const isuzu = "Isuzu"
	descrip := Isuzu{}

	groupj := NewWMIGroup("J")
	groupj.Add("AA", isuzu, MPV, descrip)
	groupj.Add("AB", isuzu, MPV, descrip)
	groupj.Add("AC", isuzu, MPV, descrip)
	groupj.Add("AL", isuzu, IncompleteCar, descrip)
	groupj.Add("81", isuzu, MPV, descrip)
	groupj.Add("82", isuzu, MPV, descrip)
}*/
