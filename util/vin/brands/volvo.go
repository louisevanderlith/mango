package brands

import "github.com/louisevanderlith/mango/util/vin/common"

type Volvo struct {
	common.VDS
}

func (v Volvo) GetPassengerCar(sections common.VINSections, year int) common.VDS {

	vehicleCode := sections.VDSCode.Char4
	platformCode := sections.VDSCode.Char5
	engineCode := sections.VDSCode.GetTypeCode("6", "7")
	bodyEmissionCode := sections.VDSCode.Char8

	dirt := Dirty{}

	findVolvoModel(vehicleCode, year, &dirt)
	findVolvoPos5(platformCode, year, &dirt)
	findVolvoEngine(engineCode, year, &dirt)
	findVolvoPos8(bodyEmissionCode, year, &dirt)
	findVolvoGearbox(sections.VISCode.CheckDigit, &dirt)

	v.VDS = compileDirty(dirt)
	v.AssemblyPlant = findVolvoAssemblyPlant(sections.VISCode.AssemblyPlantCode)

	return v.VDS
}

func findVolvoModel(typeCode string, year int, dirt *Dirty) {
	types := make(map[string]Data)

	if year <= 2006 {
		types["A"] = Data{Model: "240"}
	} else if year <= 2015 {
		types["A"] = Data{Model: "S80"}
	} else {
		types["A"] = Data{Model: "XC90"}
	}

	if year <= 2008 {
		types["B"] = Data{Model: "260"}
	} else {
		types["B"] = Data{Model: "V70, XC70"}
	}

	types["C"] = Data{Model: "XC90"}
	types["D"] = Data{Model: "XC60"}

	if year <= 2010 {
		types["F"] = Data{Model: "740"}
		types["G"] = Data{Model: "760"}
	} else {
		types["F"] = Data{Model: "S60, V60"}
		types["G"] = Data{Model: "V60", Extra: "Plug-In Hybrid"}
	}

	types["H"] = Data{Model: "780"}
	types["J"] = Data{Model: "V70"}
	types["K"] = Data{Model: "960"}

	if year <= 1998 {
		types["L"] = Data{Model: "850, S70"}
	} else {
		types["L"] = Data{Model: "V70"}
	}

	if year <= 2013 {
		types["M"] = Data{Model: "S40, V50, C30, C70"}
		types["R"] = Data{Model: "S40, V50, C30, C70"}
	} else {
		types["M"] = Data{Model: "V40"}
		types["R"] = Data{Model: "S60"}
	}

	types["S"] = Data{Model: "V70, XC70"}
	types["T"] = Data{Model: "S80"}
	types["V"] = Data{Model: "V40"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func findVolvoPos5(typeCode string, year int, dirt *Dirty) {
	if year <= 1991 {
		findVolvoSafety(typeCode, dirt)
	} else if year <= 1998 {
		findVolvoBodynSafety(typeCode, dirt)
	} else {
		findVolvoPlatform(typeCode, year, dirt)
	}
}

func findVolvoSafety(typeCode string, dirt *Dirty) {
	types := make(map[string]Data)

	types["A"] = Data{Extra: "Air Bag, 3-Point Safety Harness (Seat Belt)"}
	types["X"] = Data{Extra: "3-Point Safety Harness (Seat Belt)"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func findVolvoBodynSafety(typeCode string, dirt *Dirty) {
	types := make(map[string]Data)

	types["S"] = Data{Doors: "4 Door", BodyStyle: "Sedan", Extra: "Air Bag, 3-Point Safety Harness (Seat Belt)"}
	types["W"] = Data{Doors: "5 Door", BodyStyle: "Wagon", Extra: "Air Bag, 3-Point Safety Harness (Seat Belt)"}
	types["T"] = Data{Doors: "4 Door", BodyStyle: "Sedan", Extra: "Air Bag, 3-Point Safety Harness (Seat Belt)"}
	types["X"] = Data{Doors: "5 Door", BodyStyle: "Wagon", Extra: "Air Bag, 3-Point Safety Harness (Seat Belt)"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func findVolvoPlatform(typeCode string, year int, dirt *Dirty) {
	types := make(map[string]Data)

	types["C"] = Data{Model: "C70"}
	types["H"] = Data{Model: "S40, S60, S80", Extra: "AWD"}
	types["J"] = Data{Model: "V50, V70", Extra: "AWD"}
	types["K"] = Data{Model: "C30", Extra: "FWD"}
	types["L"] = Data{Model: "XC60", Extra: "2WD"}

	if year <= 2013 {
		types["M"] = Data{Model: "XC90", Extra: "AWD, 5 Seater"}
	} else {
		types["M"] = Data{Model: "V40", Extra: "Cross Country, AWD"}
	}

	types["N"] = Data{Model: "XC90", Extra: "5 Seater, FWD"}
	types["R"] = Data{Model: "XC90", Extra: "5 Seater, AWD"}
	types["S"] = Data{Model: "S40, S60, S80", Extra: "FWD"}
	types["V"] = Data{Model: "V40", Extra: "FWD"}
	types["W"] = Data{Model: "V50, V60, V70", Extra: "FWD, AWD, Plug-In Hybrid"}
	types["Y"] = Data{Model: "XC90", Extra: "7 Seater, FWD"}
	types["Z"] = Data{Model: "XC60, XC70, XC90", Extra: "AWD, 7 Seater"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func findVolvoEngine(typeCode string, year int, dirt *Dirty) {
	types := make(map[string]Data)

	types["17"] = Data{Model: "V40", EngineModel: "B4204S2", EngineSize: "2.0L", Extra: "FWD"}
	types["18"] = Data{Model: "V40", EngineModel: "B4194T", EngineSize: "1.9L", Extra: "Turbo, FWD"}
	types["20"] = Data{Model: "C30", EngineModel: "B4164S3", EngineSize: "1.6L", Extra: "FWD"}
	types["21"] = Data{Model: "V50, S40", EngineModel: "B4184S11", EngineSize: "1.8L"}
	types["30"] = Data{Model: "XC90", EngineModel: "D5244T18", EngineSize: "2.4L", Extra: "AWD"}
	types["38"] = Data{Model: "S40, V50", EngineModel: "B5244S4", EngineSize: "2.4iL", Extra: "FWD"}
	types["39"] = Data{Model: "S40, V50", EngineModel: "B5244S7", EngineSize: "2.4iL", Extra: "FWD"}
	types["40"] = Data{Model: "S60, V60", EngineModel: "B4204T11", EngineSize: "2.0L", Extra: "T5, FWD"}
	types["41"] = Data{Model: "850, V70", EngineModel: "B5202S", EngineSize: "2.0iL", Extra: "FWD"}
	types["43"] = Data{Model: "850, V70, S80", EngineModel: "B5204T3", EngineSize: "2.0L", Extra: "T5, FWD"}
	types["47"] = Data{Model: "850, V70", EngineModel: "B5204T", EngineSize: "2.0L", Extra: "T5, FWD"}
	types["51"] = Data{Model: "850, V70, V40, V40CC", EngineModel: "B5252S, D5204T6", EngineSize: "2.5iL", Extra: "FWD, D3, D4"} // B5252S 850/V70 2.5i FWD or D5204T6 V40/V40CC D3/D4'
	types["52"] = Data{Model: "S60, V70", EngineModel: "B5254T4", Extra: "R, AWD"}
	types["53"] = Data{Model: "S60, V70", EngineModel: "B5234T3", Extra: "T5, FWD"}
	types["54"] = Data{Model: "S60", EngineModel: "B5244T5", Extra: "T5, FWD"}
	types["55"] = Data{Model: "850, V70, S70", EngineModel: "B5254FS", Extra: "FWD"}

	if year <= 1999 {
		types["56"] = Data{Model: "S70, V70", EngineModel: "B5254T", Extra: "GLT, FWD"}
	} else {
		types["56"] = Data{EngineModel: "B5244T"}
	}

	types["57"] = Data{Model: "850, V70", EngineModel: "B5234T", Extra: "FWD, Turbo"}

	if year >= 1995 && year <= 1997 {
		types["58"] = Data{Model: "850", EngineModel: "B5234T5", Extra: "T5R, R"}
	} else if year >= 2000 {
		types["58"] = Data{Model: "S60, V70, S80, XC70", EngineModel: "B5244T3"}
	}

	types["59"] = Data{Model: "S80, S60, XC90, V70, XC70", EngineModel: "B5254T2", EngineSize: "2.5L", Extra: "Turbo, FWD, AWD"}
	types["61"] = Data{Model: "S60, V70", EngineModel: "B5244S", EngineSize: "2.4L", Extra: "FWD"}
	types["64"] = Data{Model: "S60, V70", EngineModel: "B5244S6", EngineSize: "2.4L", Extra: "FWD"}
	types["65"] = Data{Model: "S80, V70", EngineModel: "B5244S2", EngineSize: "2.4L", Extra: "FWD"}
	types["66"] = Data{Model: "S40, V50", EngineModel: "B5244S5", EngineSize: "2.4L", Extra: "FWD"}
	types["67"] = Data{Model: "C30, C70", EngineModel: "B5244S4", EngineSize: "2.4L", Extra: "T5, FWD"}
	types["68"] = Data{Model: "S40, V50", EngineModel: "B5254T3", Extra: "T5, FWD, AWD"}
	types["69"] = Data{Model: "S80, V70", EngineModel: "D5244T5", EngineSize: "2.4D", Extra: "Diesel, FWD"}
	types["70"] = Data{Model: "S40, V40, XC60", EngineModel: "D4192T3, D5244T10", Extra: "D5(205), AWD"} // D4192T3 S40/V40 / D5244T10 XC60 AWD D5(205)
	types["71"] = Data{Model: "V70, XC90", EngineModel: "D5244T4", Extra: "AWD, D5(185)"}
	types["72"] = Data{Model: "S70, S80", EngineModel: "D5252T", EngineSize: "2.5TDi", Extra: "TDi, FWD"}
	types["73"] = Data{Model: "S40, V40", EngineModel: "D4192T2"}
	types["74"] = Data{EngineModel: "D5244T2"}
	types["75"] = Data{Model: "C30", EngineModel: "D4204T", EngineSize: "2.0D", Extra: "Diesel"}
	types["76"] = Data{EngineModel: "D4164T", EngineSize: "1.6D", Extra: "PSA-Ford Engine"}

	if year >= 1981 && year <= 1984 {
		types["77"] = Data{EngineModel: "D24"}
	} else {
		types["77"] = Data{Model: "S40", EngineModel: "D5244T8", Extra: "D5", Transmission: "Automatic"}
	}

	types["78"] = Data{Model: "S40, V40", EngineModel: "D4192T4"}
	types["79"] = Data{EngineModel: "D5244T", Extra: "D5(163)"}
	types["82"] = Data{Model: "XC60", EngineModel: "D5244T15", Extra: "AWD, D5(215)"}

	types["84"] = Data{EngineModel: "D4162T", Model: "S60"}
	types["85"] = Data{EngineModel: "B8444S", Model: "XC90, S80", Extra: "V8, AWD"}

	if year >= 1983 && year <= 1984 {
		types["88"] = Data{EngineModel: "B23F"}
	} else if year >= 1985 && year <= 1995 {
		types["88"] = Data{EngineModel: "B230F"}
	} else {
		types["88"] = Data{EngineModel: "D5204T3", Model: "XC60"}
	}

	types["90"] = Data{EngineModel: "B6284T", Model: "S80", EngineSize: "2.8", Extra: "T6"}
	types["91"] = Data{EngineModel: "B6294T", Model: "S80, XC90", EngineSize: "2.9", Extra: "T6"}
	types["94"] = Data{EngineModel: "B6294S", Model: "S80", EngineSize: "2.9", Extra: "FWD"}
	types["97"] = Data{EngineModel: "B6299S", Model: "S80", EngineSize: "2.9", Extra: "FWD"}
	types["98"] = Data{EngineModel: "B6324S", Model: "XC90, S80, V70, XC70", EngineSize: "3.2", Extra: "FWD, AWD"}
	types["99"] = Data{EngineModel: "B6304T4", Model: "S80", EngineSize: "3.0", Extra: "T6, AWD"}
	types["AA"] = Data{Model: "V60", Extra: "Plug-In Hybrid"}
	types["11"] = Data{EngineModel: "B17A"}
	types["24"] = Data{EngineModel: "B19E"}
	types["26"] = Data{EngineModel: "B19ET"}
	types["37"] = Data{EngineModel: "D20"}
	types["44"] = Data{EngineModel: "B21E"}
	types["45"] = Data{EngineModel: "B21F"}
	types["46"] = Data{EngineModel: "B21ET"}
	types["48"] = Data{EngineModel: "B21LH", Extra: "Jet"}
	types["49"] = Data{EngineModel: "B21F", Extra: "MPG (1981), Kjet (1982)"}
	types["62"] = Data{EngineModel: "B28A"}
	types["80"] = Data{EngineModel: "B230G"}
	types["81"] = Data{EngineModel: "B23A"}
	types["83"] = Data{EngineModel: "B230FD", Extra: "w/ Pulsair"}
	types["86"] = Data{EngineModel: "B230FT", Extra: "EGR"}
	types["87"] = Data{EngineModel: "B230FT", Extra: "non EGR"}
	types["89"] = Data{EngineModel: "B234F"}
	types["92"] = Data{EngineModel: "B6244F"}
	types["93"] = Data{EngineModel: "B6254F"}
	types["95"] = Data{EngineModel: "B6304F", Extra: "with air pump"}
	types["96"] = Data{EngineModel: "B6204F", Extra: "without air pump"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func findVolvoPos8(typeCode string, year int, dirt *Dirty) {
	if year <= 1991 {
		findVolvoSafetyEquipment(typeCode, dirt)
	} else {
		findVolvoECE(typeCode, dirt)
	}
}

func findVolvoSafetyEquipment(typeCode string, dirt *Dirty) {
	types := make(map[string]Data)

	types["2"] = Data{Doors: "2 Door"}
	types["4"] = Data{Doors: "4 Door", BodyStyle: "Sedan"}
	types["5"] = Data{Doors: "5 Door", BodyStyle: "Wagon"}
	types["7"] = Data{Doors: "2 Door", BodyStyle: "Coupe, Bertone"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func findVolvoECE(typeCode string, dirt *Dirty) {
	types := make(map[string]Data)

	types["0"] = Data{Extra: "SULEV+", EngineModel: "B5244S7, B5254FS, B5244S6, D5252T"}
	types["2"] = Data{Extra: "ULEV2", EngineModel: "B5244S4, B5202S, B5252S, D5204T6, B5254T2, B5244S, B5244S4, B5254T3, B8444S, B6324S, B6304T4"}
	types["4"] = Data{EngineModel: "D5244T4"}
	types["7"] = Data{Extra: "LEV2", EngineModel: "B5254T4, B5244T5"}
	types["8"] = Data{EngineModel: "D4192T3, D5244T10"}
	types["D"] = Data{Extra: "L6"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func findVolvoGearbox(typeCode string, dirt *Dirty) {
	types := make(map[string]Data)

	types["1"] = Data{Transmission: "Manual", Extra: "M90, M56"}
	types["2"] = Data{Transmission: "Manual", Extra: "M46, M59, MTX75"}
	types["3"] = Data{Transmission: "Manual", Extra: "M47, M58"}
	types["4"] = Data{Transmission: "Manual", Extra: "M66"}
	types["5"] = Data{Extra: "ZF22"}
	types["6"] = Data{Transmission: "Automatic", Extra: "AW70 Lock-up, AW71 Lock-up, AW72 Lock-up, AW42"}
	types["7"] = Data{Transmission: "Automatic", Extra: "AW70 no Lock-up, AW71 no Lock-up, AW50AWD"}
	types["8"] = Data{Transmission: "Automatic", Extra: "AW42AWD"}
	types["9"] = Data{Transmission: "Automatic", Extra: "AW55-50SN, AW55-51, S40D5"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func findVolvoAssemblyPlant(typeCode string) string {
	types := make(map[string]string)

	types["0"] = "Sweden, Kalmar Plant"
	types["1"] = "Sweden, Torslanda Plant VCT 21(Volvo Torslandaverken) (Gothenburg)"
	types["2"] = "Belgium, Ghent Plant VCG 22"
	types["3"] = "Canada, Halifax Plant"
	types["4"] = "Italy, Bertone models 240"
	types["5"] = "Malaysia"
	types["6"] = "Australia"
	types["7"] = "Indonesia"
	types["A"] = "Sweden, Uddevalla Plant (Volvo Cars/TWR (Tom Walkinshaw Racing))"
	types["B"] = "Italy, Bertone Chongq 31"
	types["D"] = "Italy, Bertone models 780"
	types["E"] = "Singapore"
	types["F"] = "The Netherlands, Born Plant (NEDCAR)"
	types["J"] = "Sweden, Uddevalla Plant, VCU 38 (Volvo Cars/ Pininfarina Sverige AB)"
	types["M"] = "PVÖ 53"

	result := "Unknown"
	if val, ok := types[typeCode]; ok {
		result = val
	}

	return result
}

/*
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
}*/
