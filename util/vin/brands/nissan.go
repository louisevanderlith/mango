package brands

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type Nissan struct {
	common.VDS
}

// https://www.nedautoparts.com/pages/resources#
// http://nissanhelp.com/diy/common/nissan_engine_decoder.php
// https://x.nissanhelp.com/forums/Knowledgebase/links/16/?catid=16
func (v Nissan) GetPassengerCar(sections common.VINSections, year int) common.VDS {
	return v.VDS
}

func decodeEngine(engineCode string) {

}

func decodeModel(modelCode string) {
	types := make(map[string]string)

	types["310"] = "Nissan Sedan/Wagon"
	types["410"] = "Nissan 410"
	types["411"] = "Nissan 411"
	types["510"] = "Nissan 510"
	types["520/1"] = "Nissan 520 (Truck)"
	types["60"] = "Nissan Patrol"
	types["610S"] = "Nissan 610 Sedan"
	types["610W"] = "Nissan 610 Wagon"
	types["620"] = "Nissan 620 Truck"
	types["710S"] = "Nissan 710 Sedan"
	types["710W"] = "Nissan 710 Wagon"
	types["720"] = "Nissan 720 Truck"
	types["810"] = "Nissan 810"
	types["910"] = "Nissan 810 Maxima"
	types["A10"] = "Nissan 510 (78-81)"
	types["A32"] = "Nissan Maxima (1995-1999)"
	types["A33"] = "Nissan Maxima (2000-2003)"
	types["A34"] = "Nissan Maxima (2004-2009)"
	types["A35"] = "Nissan Maxima (2010-2012)"
	types["A60"] = "Nissan Titan (2003-2012)"
	types["B11"] = "Nissan Sentra (83-86)"
	types["B110"] = "Nissan 1200"
	types["B12"] = "Nissan Sentra (1987-1990)"
	types["B13"] = "Nissan Sentra (1991-1994)"
	types["B14"] = "Nissan Sentra (1995-1999)"
	types["B15"] = "Nissan Sentra (2000-2006)"
	types["B16"] = "Nissan Sentra (2006-2012)"
	types["B210"] = "Nissan B210"
	types["B310"] = "Nissan 210"
	types["C11"] = "Nissan Versa (2007-2012)"
	types["CL32"] = "Nissan Altima Coupe (2008-2012)"
	types["D21"] = "Nissan Hardbody Truck (86.5-97)"
	types["D22"] = "Nissan Frontier Truck (98~)"
	types["E52"] = "Nissan Quest (2011-2012))"
	types["F10"] = "Nissan F10"
	types["F15"] = "Nissan JUKE (2011-2012)"
	types["F80"] = "Nissan NV"
	types["GC22"] = "Nissan Van"
	types["HL32"] = "Nissan Altima Hybrid (2007-2011)"
	types["J30"] = "Nissan Maxima (89-94)"
	types["KN13"] = "Nissan Pulsar 87-90)"
	types["L30"] = "Nissan Alt (98-01)"
	types["L31"] = "Nissan Altima (02~)"
	types["M10"] = "Nissan Stanza Wagon"
	types["M11"] = "Nissan Axxess"
	types["N10"] = "Nissan 310"
	types["N12"] = "Nissan Pulsar (83-86)"
	types["N50"] = "Nissan Nissan Xterra (2005-2012)"
	types["R35"] = "Nissan GTR (2009-2012)"
	types["R50"] = "Nissan Pathfinder (1996-2004)"
	types["R51"] = "Nissan Pathfinder (2005-2012)"
	types["S10"] = "Nissan 200SX (77-79)"
	types["S110"] = "Nissan 200SX (80-83)"
	types["S12"] = "Nissan 200SX (84-88)"
	types["S13"] = "Nissan 240SX (89-94)"
	types["S130"] = "Nissan 280ZX (79-83)"
	types["S14"] = "Nissan 240SX (95-98)"
	types["S30"] = "Nissan 240Z, 260Z, 280Z, (70-78)"
	types["S35"] = "Nissan Rogue(2008-2012)"
	types["SPSR"] = "Nissan Roadster"
	types["T11"] = "Nissan Stanza"
	types["T12"] = "Nissan Stanza"
	types["TA60"] = "Nissan Armada (2004-2012)"
	types["U11"] = "Nissan Maxima (86-88)"
	types["U12"] = "Nissan Stanza (90-92)"
	types["U13"] = "Nissan Altima (93-97)"
	types["V40"] = "Nissan Quest (93-98)"
	types["V41"] = "Nissan Quest (2000-2002))"
	types["V42"] = "Nissan Quest (2004-2010))"
	types["WD21"] = "Nissan Pathfinder (86-95)"
	types["WD22"] = "Nissan Xterra (2000-2004)"
	types["Z12"] = "Nissan CUBE (2009-2012)"
	types["Z31"] = "Nissan 300ZX (84-89)"
	types["Z32"] = "Nissan 300ZX (90-96)"
	types["Z50"] = "Nissan Murano (2003-2007)"
	types["Z51"] = "Nissan Murano (2009-2012)"

	return types[modelCode]
}
