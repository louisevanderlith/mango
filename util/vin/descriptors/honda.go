package descriptors

import (
	"fmt"

	"github.com/louisevanderlith/mango/util/vin"
)

// 1: Country of origin
// 2: Manufacturer
// 3: Make
// 4-6: Chassis and engine
// 7: Body and transmission
// 8: Trim level and restraint
// 9: Check digit
// 10: Model year
// 11: Assembly plant
// 12-17: Production number
type Honda struct {
	Model          string
	BodyType       string
	Engine         string
	Transmission   string
	SequenceNumber string
}

// GetData Deserializes Honda VIN Numbers
func (d Honda) GetData(vinNo string) string {
	yearCode := vinNo[9:10]
	years := vin.GetYear(yearCode)
	fmt.Println("Code:", yearCode, "Years:", years)

	modelCode := vinNo[3:6]
	model := modelTypes(modelCode, years[0])
	fmt.Println("Code:", modelCode, "Model:", model)

	bodyCode := vinNo[3:5]
	bodyType := bodyTypes(bodyCode)
	fmt.Println("Code:", bodyCode, "Body:", bodyType)

	transCode := vinNo[6:7]
	trans := transTypes(transCode)
	fmt.Println("Code:", transCode, "Trans:", trans)

	bodynTransCode := vinNo[6:7]
	bodynTrans := bodynTransTypes(bodynTransCode, years[0])
	fmt.Println("Code:", bodynTransCode, "Body n Trans:", bodynTrans)

	gradeCode := vinNo[8:9]
	grade := gradeTypes(gradeCode, years[0], model)
	fmt.Println("Code:", gradeCode, "Grade:", grade)

	seq := vinNo[13:]
	fmt.Println("SEQ:", seq)

	result := ""
	result += fmt.Sprintf("Year: %s \r\n")
	result += fmt.Sprintf("Year: %s \r\n")
	result += fmt.Sprintf("Year: %s \r\n")
	result += fmt.Sprintf("Year: %s \r\n")

	return result
}

// String returns a text representation of this descriptor
func (d Honda) String() string {
	return ""
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

	result := "Unknown"

	if val, ok := types[typeCode]; ok {
		result = val
	}

	return result
}

func transTypes(typeCode string) string {
	types := make(map[string]string)
	types["2"] = "Semi-Hondamatic"
	types["3"] = "3 speed automatic"
	types["4"] = "4 speed manual"
	types["5"] = "5 speed manual"
	types["6"] = "5 speed manual super-low gear"
	types["7"] = "4 speed automatic"
	types["8"] = "5 speed automatic"

	result := "Unknown"

	if val, ok := types[typeCode]; ok {
		result = val
	}

	return result
}

func modelTypes(typeCode string, year int) string {
	types := make(map[string]string)

	if year <= 1989 {
		types["BA3"] = "Prelude A20 Engine"
		types["BA4"] = "Prelude B20 Engine"
		types["BA6"] = "Prelude A18 Engine"
		types["CA5"] = "Accord"
		types["CA6"] = "Accord coupe"
		types["EC1"] = "Civic CRX"
		types["EC2"] = "Civic 1300"
		types["EC3"] = "Civic 1500"
		types["EC4"] = "Civic 4 Door Sedan"
		types["EC5"] = "Civic Wagon"
		types["EC6"] = "Civic Wagon 4X4"
		types["ED3"] = "Civic Sedan 1.5L (1988)"
		types["ED6"] = "Civic Hatchback 1.5L (1988)"
		types["ED7"] = "Civic Hatchback 1.6L (1988)"
		types["ED8"] = "Civic CRX 1.5L (1988)"
		types["ED9"] = "Civic CRX 1.6L (1988)"
		types["EE2"] = "Civic Wagon 1.5L (1988)"
		types["EE4"] = "Civic Wagon 1.6L (1988)"
		types["EY1"] = "Civic Wagovan"
		types["EY3"] = "Civic Wagovan 1.5L"
	} else {
		types["BA4"] = "Prelude, 2.0/2.1L"
		types["BA8"] = "Prelude, 2.2L"
		types["BB"] = "Prelude VTEC, 2.2L"
		types["BB2"] = "Prelude, 2.3L"
		types["BB6"] = "Prelude 2-door"
		types["CB3"] = "Japan Version Accord, 2.2L"
		types["CB7"] = "Accord, 2.2L"
		types["CB9"] = "Accord Wagon, 2.2L"
		types["CC1"] = "Accord coupé, 2.0L"
		types["CD5"] = "Accord 4 Door 2.2L VTEC"
		types["CD7"] = "Accord 2 Door 2.2L"
		types["CE1"] = "Accord Wagon 2.2L"
		types["CE6"] = "Accord V6 4 Door 2.7L"
		types["CF8"] = "Accord 4 Door SOHC"
		types["CF4"] = "Accord 4 Door DOHC SiR-T"
		types["CG1"] = "Accord 4 Door V6 VTEC"
		types["CG2"] = "Accord 2 Door V6 VTEC"
		types["CG3"] = "Accord 2 Door (VTEC or ULEV)"
		types["CG5"] = "Accord 4 Door VTEC"
		types["CG6"] = "Accord 4 Door ULEV"
		types["CH9"] = "Accord Wagon SiR 2.3L (VTEC)"
		types["ED3"] = "Civic 4 Door, 1.5L"
		types["ED4"] = "Civic 4 Door, 1.6L"
		types["ED6"] = "Civic 3 Door, 1.5L"
		types["ED7"] = "Civic3 Door, 1.6L"
		types["ED8"] = "CRX, 1.5L"
		types["ED9"] = "CRX, 1.6L"
		types["EE2"] = "Civic Wagon, 1.5L"
		types["EE4"] = "Civic Wagon, 1.6L"
		types["EE8"] = "CR-X 1.6L VTEC"
		types["EM1"] = "Civic Si 1.6L VTEC"
		types["EG1"] = "Civic del Sol, 1.5L"
		types["EG2"] = "Civic del Sol VTEC 1.6L"
		types["EG6"] = "Civic 3 door hatchback 1.6L"
		types["EG8"] = "Civic 4 Door, 1.5L"
		types["EG9"] = "Civic 4 Door, 1.6L VTEC"
		types["EH2"] = "Civic 3 Door, 1.5L"
		types["EH3"] = "Civic 3 Door, 1.6L"
		types["EH6"] = "Civic del Sol, 1.6L"
		types["EH9"] = "Civic 4 Door, 1.6L"
		types["EJ1"] = "Civic2 Door, 1.6L"
		types["EJ2"] = "Civic2 Door, 1.5L"
		types["EJ6"] = "Civic 2/3/4 Door 1.6L"
		types["EJ7"] = "Civic 2 Door 1.6L"
		types["EJ8"] = "Civic 2/4 Door 1.6L"
		types["EJ9"] = "Civic 3 Door 1.4L"
		types["EL1"] = "Orthia 1.8L 5 Door Wagon"
		types["EL2"] = "Orthia 2.0L 5 Door Wagon"
		types["EL3"] = "Orthia 2.0L 5 Door Wagon 4WD"
		types["RA1"] = "Odyssey"
		types["RA3"] = "Odyssey 5 Door Wagon (1998)"
		types["RD1"] = "CR-V 5-door (4 Wheel drive)"
		types["RD2"] = "CR-V 5-door (2 Wheel drive)"
	}

	result := "Unknown"

	if val, ok := types[typeCode]; ok {
		result = val
	}

	return result
}

func bodynTransTypes(typeCode string, year int) string {
	types := make(map[string]string)

	if year <= 1986 {
		types["2"] = "2 door"
		types["3"] = "3 door"
		types["4"] = "4 door"
		types["5"] = "5 door"
	} else if year >= 1987 && year <= 1989 {
		types["1"] = "2 Door Sedan, Manual"
		types["2"] = "2 Door Sedan, Manual"
		types["3"] = "2 Door Hatchback, Manual"
		types["4"] = "2 Door Hatchback, Automatic"
		types["5"] = "4 Door Sedan, Manual"
		types["6"] = "4 Door Sedan, Automatic"
		types["7"] = "4 Door Wagon, Manual"
		types["8"] = "4 Door Wagon, Automatic"
	} else if year >= 1990 {
		types["1"] = "2 Door Coupe, Manual"
		types["2"] = "2 Door Coupe, Automatic"
		types["3"] = "3 Door Hatchback, Manual"
		types["4"] = "3 Door Hatchback, Automatic"
		types["5"] = "4 Door Sedan, Manual"
		types["6"] = "4 Door Sedan, Automatic"
		types["7"] = "5 Door Wagon, Manual"
		types["8"] = "5 Door Wagon, Automatic"
	}

	result := "Unknown"

	if val, ok := types[typeCode]; ok {
		result = val
	}

	return result
}

func gradeTypes(typeCode string, year int, model string) string {
	types := make(map[string]string)

	if year <= 1987 {
		types["1"] = "Basic, HF"
		types["2"] = "Standard, DX"
		types["3"] = "GL, GLS, LX"
		types["4"] = "LXi, Si"
		types["5"] = "Special"
		types["6"] = "Accord Hatchback"
		types["7"] = "Accord Hatchback & LXi"
		types["8"] = "Accord Hatchback & LXi ('passive') seat belt"
	} else if year >= 1988 && year <= 1989 {
		switch model {
		case "Accord":
			types["2"] = "DX Accord Manual seat belt"
			types["3"] = "LX Accord Manual seat belt"
			types["4"] = "LXi Accord Manual seat belt"
			types["5"] = "SEi Accord Manual seat belt"
			types["6"] = "DX Accord Automatic seat belt"
			types["8"] = "LXi Accord Automatic ('passive') seat belt"
		case "Prelude":
			types["2"] = "S Prelude Automatic ('passive') seat belt"
			types["3"] = "Si Prelude Automatic ('passive') seat belt"
			types["4"] = "Si Prelude w/optional 4WS Automatic ('passive') seat belt"
		case "Civic":
			types["4"] = "(std) Civic Hatchback Manual seat belt | DX Civic Sedan Manual seat belt | (std) Civic Wagovan Manual seat belt"
			types["5"] = "DX Civic Hatchback Manual seat belt | LX Civic Sedan Manual seat belt | (std) Civic CRX Manual seat belt | (std) '89 Civic CRX Automatic ('passive') seat belt | (std) Civic Wagon Manual seat belt"
			types["6"] = "HF Civic CRX Manual seat belt | Si Civic CRX Manual seat belt | Si Civic CRX Automatic ('passive') seat belt | Si Civic Hatchback Manual seat belt | RTi Civic Wagon 4WD Manual seat belt"
		}
	} else if year >= 1990 && year <= 1993 {
		switch model {
		case "Accord":
			types["4"] = "DX Accord"
			types["5"] = "LX Accord"
			types["6"] = "DC Accord"
			types["7"] = "EX Accord (92-93)"
			types["8"] = "SE Accord"
			types["9"] = "Accord 10th Anniversary"
		case "Prelude":
			types["1"] = "2.0 S Prelude"
			types["2"] = "2.0 S Prelude"
			types["3"] = "2.1 S Prelude"
			types["4"] = "2.1 Si Prelude w/4WS | S Prelude (92-93)"
			types["5"] = "2.3 Si Prelude"
			types["6"] = "Si Prelude w/4WS (92-93)"
			types["7"] = "VTEC Prelude (1993)"
		case "Civic":
			types["4"] = "(std) Civic 3 Door | CX Civic 3 Door (1992) | DX Civic | S Civic del Sol (1993)"
			types["5"] = "LX Civic | CX Civic (1993) | DX Civic 3 Door | DX Civic Wagon (2WD)"
			types["6"] = "DX Civic 3 Door (1993) | EX Civic 4 Door | EX Civic 2 Door (1993) | Si Civic 3 Door | Si Civic del Sol (1993) | ATi Civic Wagon (4WD) | VX Civic 3 Door (1992)"
			types["7"] = "VX Civic (1993)"
			types["8"] = "Si Civic 3 Door (92-93)"
		case "CRX":
			types["6"] = "HF CRX | Si CRX"
		}
	} else if year >= 1994 && year <= 1999 {
		switch model {
		case "Accord":
			types["0"] = "Accord SE 2/4 Door"
			types["1"] = "Accord DX w/ABS 2/4 Door"
			types["2"] = "Accord DX 2/4 Door | Accord Wagon LX"
			types["3"] = "Accord LX 2/4 Door | Accord Wagon LX w/ABS | Accord w/ABS V6"
			types["4"] = "Accord DX 4 Door (1998) | Accord LX w/ABS 2/4 Door"
			types["5"] = "Accord EX 2/4 Door"
			types["6"] = "Accord E w/leather 2/4 Door"
			types["7"] = "Accord EX | Accord EX ULEV (1998)"
			types["8"] = "Accord DX 25th Anniversary"
			types["9"] = "Accord Wagon EX | Accord Value Package 4 Door"
		case "Prelude":
			types["4"] = "Prelude S | Prelude VTEC (1997) | Prelude w/o SH pkg. (1998)"
			types["5"] = "Prelude Si | Prelude SH (1997)"
			types["6"] = "Prelude 4WS Si"
			types["7"] = "Prelude VTEC"
		case "Civic":
			types["0"] = "Civic LX"
			types["2"] = "Civic CX 3 Door | Civic DX 2/4 Door | Civic EX 2 Door | Civic HX | Civic CVT"
			types["3"] = "Civic CX w/AC 3 Door | Civic DX w/AC 4 Door | Civic EX w/ABS 2 Door | Civic LX"
			types["4"] = "Civic DX 4 Door | Civic EX 2/4 Door | Civic del Sol S"
			types["5"] = "Civic CX 3 Door | Civic LX 4 Door | Civic DX w/AC | Civic EX a/ABS 2 Door"
			types["6"] = "Civic DX 3 Door | Civic LX w/ABS 4 Door | Civic del Sol Si"
			types["7"] = "Civic LX w/AC | Civic del Sol Si w/ABS | Civic del Sol VTEC | Civic VX 3 Door"
			types["8"] = "Civic Si 3 Door | Civic LX w/ABS"
			types["9"] = "Civic del Sol LX w/ABS | Civic del Sol Si w/ABS | Civic del Sol VTEC w/ABS | Civic Si w/ABS 3 Door | Civic EX 4 Door"
		case "CR-V":
			types["4"] = "CR-V w/o ABS | CR-V LX"
			types["5"] = "CR-V w/ABS"
			types["6"] = "CR-V EX (98-99)"
		case "Odyssey":
			types["4"] = "Odyssey LX 6 Passaway"
			types["6"] = "Odyssey LX 7 Passenger"
			types["7"] = "Odyssey EX"
		}
	}

	result := "Unknown"

	if val, ok := types[typeCode]; ok {
		result = val
	}

	return result
}

func groupsHonda() {
	const honda = "Honda"
	const acura = "Acura"
	descrip := Honda{}

	groupj := NewWMIGroup("J")
	groupj.Add("D0", honda, Motorcycle, descrip)
	groupj.Add("HF", honda, NotSpecified, descrip)
	groupj.Add("HG", honda, NotSpecified, descrip)
	groupj.Add("HM", honda, PassengerCar, descrip)
	groupj.Add("HL", honda, MPV, descrip)
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
	groupp.Add("AD", honda, MPV, descrip)
	groupp.Add("MH", honda, PassengerCar, descrip)
	groups := NewWMIGroup("S")
	groups.Add("HH", honda, PassengerCar, descrip)
	groups.Add("HS", honda, MPV, descrip)

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
	group2.Add("HK", honda, MPV, descrip)
	group2.Add("HJ", honda, Truck, descrip)
	group2.Add("HN", acura, MPV, descrip)
	group2.Add("HU", acura, Truck, descrip)

	group3 := NewWMIGroup("3")
	group3.Add("HG", honda, PassengerCar, descrip)
	group3.Add("H1", honda, Motorcycle, descrip)

	group4 := NewWMIGroup("4")
	group4.Add("78", honda, ATV, descrip)

	group5 := NewWMIGroup("5")
	group5.Add("J6", honda, MPV, descrip)
	group5.Add("J7", honda, Truck, descrip)
	group5.Add("J8", acura, MPV, descrip)
	group5.Add("J0", acura, Truck, descrip)
	group5.Add("KB", honda, PassengerCar, descrip)
	group5.Add("KC", acura, PassengerCar, descrip)
	group5.Add("FN", honda, MPV, descrip)
	group5.Add("FP", honda, Truck, descrip)
	group5.Add("FR", acura, MPV, descrip)
	group5.Add("FS", acura, Truck, descrip)

	group9 := NewWMIGroup("9")
	group9.Add("3H", honda, PassengerCar, descrip)
	group9.Add("C2", honda, Motorcycle, descrip)
}
