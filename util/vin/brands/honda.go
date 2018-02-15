package brands

import (
	"fmt"

	"github.com/louisevanderlith/mango/util/vin/common"
)

type Honda struct {
	common.VDS
}

func (v Honda) GetPassengerCar(sections common.VINSections, year int) common.VDS {
	modelCode := sections.VDSCode.GetTypeCode("4", "5", "6", "7")
	gradeCode := sections.VDSCode.Char8

	dirt := Dirty{}

	findHondaModelTransBody(modelCode, year, &dirt)
	findHondaGradeType(gradeCode, year, cleanModel(dirt.Models), &dirt)

	v.VDS = compileDirty(dirt)

	v.AssemblyPlant = findHondaAssemblyPlant(sections.VISCode.AssemblyPlantCode)

	return v.VDS
}

// VIN 4, 5, 6, 7
func findHondaModelTransBody(modelCode string, year int, dirt *Dirty) {
	modelTypeCode := modelCode[:2]
	transTypeCode := modelCode[2:3]
	bodyAndTransCode := modelCode[3:]

	if year <= 1983 {
		modelTypePre83(modelTypeCode, dirt)
	}

	if year >= 1984 && year <= 1986 {
		modelTypePre86(modelTypeCode, dirt)
		transmissionTypePre86(transTypeCode, dirt)
		bodyTypePre86(bodyAndTransCode, dirt)
	}

	if year >= 1987 && year <= 1989 {
		modelTypePre89(modelCode[:3], year, dirt)
		bodyTypePre89(bodyAndTransCode, dirt)
	}

	if year >= 1990 && year <= 1999 {
		modelTypePre99(modelCode[:3], year, dirt)
		bodyTypePre99(bodyAndTransCode, dirt)
	}
}

func modelTypePre83(typeCode string, dirt *Dirty) {
	types := make(map[string]Data)

	types["SR"] = Data{BodyStyle: "Sedan"}
	types["WR"] = Data{BodyStyle: "Wagon"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func modelTypePre86(typeCode string, dirt *Dirty) {
	types := make(map[string]Data)

	types["AB"] = Data{Model: "Prelude"}
	types["AD"] = Data{Model: "Accord"}
	types["AE"] = Data{Model: "CRX", EngineSize: "1.3L"}
	types["AF"] = Data{Model: "CRX", EngineSize: "1.5L"}
	types["AG"] = Data{Model: "Civic", EngineSize: "1.3L", Doors: "3 Door"}
	types["AH"] = Data{Model: "Civic", EngineSize: "1.5L", Doors: "3 Door"}
	types["AK"] = Data{Model: "Civic", EngineSize: "1.5L", Doors: "4 Door"}
	types["AN"] = Data{Model: "Civic", BodyStyle: "Wagon"}
	types["AR"] = Data{Model: "Civic", BodyStyle: "Wagon", Extra: "4x4"}
	types["BA"] = Data{Model: "Accord"}
	types["AM"] = Data{Model: "Accord", EngineSize: "1.5L", Doors: "5 Door"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func modelTypePre89(typeCode string, year int, dirt *Dirty) {
	types := make(map[string]Data)

	types["BA3"] = Data{Model: "Prelude", EngineModel: "A20"}
	types["BA4"] = Data{Model: "Prelude", EngineModel: "B20"}
	types["BA6"] = Data{Model: "Prelude", EngineModel: "A18"}
	types["CA5"] = Data{Model: "Accord"}
	types["CA6"] = Data{Model: "Accord", BodyStyle: "Coupe"}
	types["EC1"] = Data{Model: "CRX"}
	types["EC2"] = Data{Model: "Civic", EngineSize: "1.3L"}
	types["EC3"] = Data{Model: "Civic", EngineSize: "1.5L"}
	types["EC4"] = Data{Model: "Civic", BodyStyle: "Sedan", Doors: "4 Doors"}
	types["EC5"] = Data{Model: "Civic", BodyStyle: "Wagon"}
	types["EC6"] = Data{Model: "Civic", BodyStyle: "Wagon", Extra: "4X4"}

	if year == 1988 {
		types["ED3"] = Data{Model: "Civic", BodyStyle: "Sedan", EngineSize: "1.5L"}
		types["ED6"] = Data{Model: "Civic", BodyStyle: "Hatchback", EngineSize: "1.5L"}
		types["ED7"] = Data{Model: "Civic", BodyStyle: "Hatchback", EngineSize: "1.6L"}
		types["ED8"] = Data{Model: "CRX", EngineSize: "1.5L"}
		types["ED9"] = Data{Model: "CRX", EngineSize: "1.6L"}
		types["EE2"] = Data{Model: "Civic", BodyStyle: "Wagon", EngineSize: "1.5L"}
		types["EE4"] = Data{Model: "Civic", BodyStyle: "Wagon", EngineSize: "1.6L"}
	}

	types["EY1"] = Data{Model: "Civic", BodyStyle: "Wagovan"}
	types["EY3"] = Data{Model: "Civic", BodyStyle: "Wagovan", EngineSize: "1.5L"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func modelTypePre99(typeCode string, year int, dirt *Dirty) {
	types := make(map[string]Data)

	types["BA4"] = Data{Model: "Prelude", EngineSize: "2.0, 2.1L"}
	types["BA8"] = Data{Model: "Prelude", EngineSize: "2.2L"}
	types["BB"] = Data{Model: "Prelude", Extra: "VTEC", EngineSize: "2.2L"}
	types["BB2"] = Data{Model: "Prelude", EngineSize: "2.3L"}
	types["BB6"] = Data{Model: "Prelude", Doors: "2 Door"}
	types["CB3"] = Data{Model: "Accord", Extra: "Japan Version", EngineSize: "2.2L"}
	types["CB7"] = Data{Model: "Accord", EngineSize: "2.2L"}
	types["CB9"] = Data{Model: "Accord", BodyStyle: "Wagon", EngineSize: "2.2L"}
	types["CC1"] = Data{Model: "Accord", BodyStyle: "Coupe", EngineSize: "2.0L"}
	types["CD5"] = Data{Model: "Accord", Doors: "4 Door", Extra: "VTEC", EngineSize: "2.2L"}
	types["CD7"] = Data{Model: "Accord", Doors: "2 Door", EngineSize: "2.2L"}
	types["CE1"] = Data{Model: "Accord", BodyStyle: "Wagon", EngineSize: "2.2L"}
	types["CE6"] = Data{Model: "Accord", Doors: "4 Door", Extra: "V6", EngineSize: "2.7L"}
	types["CF8"] = Data{Model: "Accord", Doors: "4 Door", Extra: "SOHC"}
	types["CF4"] = Data{Model: "Accord", Doors: "4 Door", Extra: "DOHC", Trim: "SiR-T"}
	types["CG1"] = Data{Model: "Accord", Doors: "4 Door", Extra: "VTEC, V6"}
	types["CG2"] = Data{Model: "Accord", Doors: "2 Door", Extra: "VTEC, V6"}
	types["CG3"] = Data{Model: "Accord", Doors: "2 Door", Extra: "VTEC, ULEV"}
	types["CG5"] = Data{Model: "Accord", Doors: "4 Door", Extra: "VTEC"}
	types["CG6"] = Data{Model: "Accord", Doors: "4 Door", Extra: "ULEV"}
	types["CH9"] = Data{Model: "Accord", BodyStyle: "Wagon", Trim: "SiR", Extra: "VTEC", EngineSize: "2.3L"}
	types["ED3"] = Data{Model: "Civic", Doors: "4 Door", EngineSize: "1.5L"}
	types["ED4"] = Data{Model: "Civic", Doors: "4 Door", EngineSize: "1.6L"}
	types["ED6"] = Data{Model: "Civic", Doors: "3 Door", EngineSize: "1.5L"}
	types["ED7"] = Data{Model: "Civic", Doors: "3 Door", EngineSize: "1.6L"}
	types["ED8"] = Data{Model: "CRX", EngineSize: "1.5L"}
	types["ED9"] = Data{Model: "CRX", EngineSize: "1.6L"}
	types["EE2"] = Data{Model: "Civic", BodyStyle: "Wagon", EngineSize: "1.5L"}
	types["EE4"] = Data{Model: "Civic", BodyStyle: "Wagon", EngineSize: "1.6L"}
	types["EE8"] = Data{Model: "CR-X", Extra: "VTEC", EngineSize: "1.6L"}
	types["EM1"] = Data{Model: "Civic", Trim: "Si", Extra: "VTEC", EngineSize: "1.6L"}
	types["EG1"] = Data{Model: "del Sol", EngineSize: "1.5L"}
	types["EG2"] = Data{Model: "del Sol", Extra: "VTEC", EngineSize: "1.6L"}
	types["EG6"] = Data{Model: "Civic", Doors: "3 Door", BodyStyle: "Hatchback", EngineSize: "1.6L"}
	types["EG8"] = Data{Model: "Civic", Doors: "4 Door", EngineSize: "1.5L"}
	types["EG9"] = Data{Model: "Civic", Doors: "4 Door", Extra: "VTEC", EngineSize: "1.6L"}
	types["EH2"] = Data{Model: "Civic", Doors: "3 Door", EngineSize: "1.5L"}
	types["EH3"] = Data{Model: "Civic", Doors: "3 Door", EngineSize: "1.6L"}
	types["EH6"] = Data{Model: "del Sol", EngineSize: "1.6L"}
	types["EH9"] = Data{Model: "Civic", Doors: "4 Door", EngineSize: "1.6L"}
	types["EJ1"] = Data{Model: "Civic", Doors: "2 Door", EngineSize: "1.6L"}
	types["EJ2"] = Data{Model: "Civic", Doors: "2 Door", EngineSize: "1.5L"}
	types["EJ6"] = Data{Model: "Civic", Doors: "2 Door, 3 Door, 4 Door", EngineSize: "1.6L"}
	types["EJ7"] = Data{Model: "Civic", Doors: "2 Door", EngineSize: "1.6L"}
	types["EJ8"] = Data{Model: "Civic", Doors: "2 Door, 4 Door", EngineSize: "1.6L"}
	types["EJ9"] = Data{Model: "Civic", Doors: "3 Door", EngineSize: "1.4L"}
	types["EL1"] = Data{Model: "Orthia", Doors: "5 Door", BodyStyle: "Wagon", EngineSize: "1.8L"}
	types["EL2"] = Data{Model: "Orthia", EngineSize: "2.0L", Doors: "5 Door", BodyStyle: "Wagon"}
	types["EL3"] = Data{Model: "Orthia", EngineSize: "2.0L", Doors: "5 Door", BodyStyle: "Wagon", Extra: "4WD"}
	types["RA1"] = Data{Model: "Odyssey"}

	if year == 1998 {
		types["RA3"] = Data{Model: "Odyssey", Doors: "5 Door", BodyStyle: "Wagon"}
	}

	types["RD1"] = Data{Model: "CR-V", Doors: "5 Door", Extra: "4WD"}
	types["RD2"] = Data{Model: "CR-V", Doors: "5 Door"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func transmissionTypePre86(transmissionCode string, dirt *Dirty) {
	types := make(map[string]Data)

	types["2"] = Data{Transmission: "Semi-Hondamatic"}
	types["3"] = Data{Gears: "3 Speed", Transmission: "Automatic"}
	types["4"] = Data{Gears: "4 Speed", Transmission: "Manual"}
	types["5"] = Data{Gears: "5 Speed", Transmission: "Manual"}
	types["6"] = Data{Gears: "5 Speed", Transmission: "Manual Super-Low Gear"}
	types["7"] = Data{Gears: "4 Speed", Transmission: "Automatic"}
	types["8"] = Data{Gears: "5 Speed", Transmission: "Automatic"}

	if val, ok := types[transmissionCode]; ok {
		fillDirty(dirt, val)
	}
}

func bodyTypePre86(typeCode string, dirt *Dirty) {
	val := fmt.Sprintf("%s Door", typeCode)
	dirt.Doors = append(dirt.Doors, val)
}

func bodyTypePre89(typeCode string, dirt *Dirty) {
	types := make(map[string]Data)

	types["1"] = Data{Doors: "2 Door", BodyStyle: "Sedan", Transmission: "Manual"}
	types["2"] = Data{Doors: "2 Door", BodyStyle: "Sedan", Transmission: "Automatic"}
	types["3"] = Data{Doors: "2 Door", BodyStyle: "Hatchback", Transmission: "Manual"}
	types["4"] = Data{Doors: "2 Door", BodyStyle: "Hatchback", Transmission: "Automatic"}
	types["5"] = Data{Doors: "4 Door", BodyStyle: "Sedan", Transmission: "Manual"}
	types["6"] = Data{Doors: "4 Door", BodyStyle: "Sedan", Transmission: "Automatic"}
	types["7"] = Data{Doors: "4 Door", BodyStyle: "Wagon", Transmission: "Manual"}
	types["8"] = Data{Doors: "4 Door", BodyStyle: "Wagon", Transmission: "Automatic"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func bodyTypePre99(typeCode string, dirt *Dirty) {
	types := make(map[string]Data)

	types["1"] = Data{Doors: "2 Door", BodyStyle: "Coupe", Transmission: "Manual"}
	types["2"] = Data{Doors: "2 Door", BodyStyle: "Coupe", Transmission: "Automatic"}
	types["3"] = Data{Doors: "3 Door", BodyStyle: "Hatchback", Transmission: "Manual"}
	types["4"] = Data{Doors: "3 Door", BodyStyle: "Hatchback", Transmission: "Automatic"}
	types["5"] = Data{Doors: "4 Door", BodyStyle: "Sedan", Transmission: "Manual"}
	types["6"] = Data{Doors: "4 Door", BodyStyle: "Sedan", Transmission: "Automatic"}
	types["7"] = Data{Doors: "5 Door", BodyStyle: "Wagon", Transmission: "Manual"}
	types["8"] = Data{Doors: "5 Door", BodyStyle: "Wagon", Transmission: "Automatic"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func findHondaGradeType(typeCode string, year int, model string, dirt *Dirty) {
	if year <= 1987 {
		gradeTypePre87(typeCode, dirt)
	} else if year >= 1988 && year <= 1989 {
		gradeTypePre89(typeCode, model, year, dirt)
	} else if year >= 1990 && year <= 1993 {
		gradeTypePre93(typeCode, model, year, dirt)
	} else if year >= 1994 && year <= 1999 {
		gradeTypePre99(typeCode, model, year, dirt)
	}
}

func gradeTypePre87(typeCode string, dirt *Dirty) {
	types := make(map[string]Data)

	types["1"] = Data{Trim: "Basic, HF"}
	types["2"] = Data{Trim: "STD, DX"}
	types["3"] = Data{Trim: "GL, GLS, LX"}
	types["4"] = Data{Trim: "LXi, Si"}
	types["5"] = Data{Trim: "Special"}
	types["6"] = Data{Model: "Accord", BodyStyle: "Hatchback"}
	types["7"] = Data{Model: "Accord", BodyStyle: "Hatchback", Trim: "LXi"}
	types["8"] = Data{Model: "Accord", BodyStyle: "Hatchback", Trim: "LXi"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func gradeTypePre89(typeCode, model string, year int, dirt *Dirty) {
	types := make(map[string]Data)

	switch model {
	case "Accord":
		types["2"] = Data{Trim: "DX", Transmission: "Manual"}
		types["3"] = Data{Trim: "LX", Transmission: "Manual"}
		types["4"] = Data{Trim: "LXi", Transmission: "Manual"}
		types["5"] = Data{Trim: "SEi", Transmission: "Manual"}
		types["6"] = Data{Trim: "DX", Transmission: "Automatic"}
		types["8"] = Data{Trim: "LXi", Transmission: "Automatic"}
	case "Prelude":
		types["2"] = Data{Trim: "S", Transmission: "Automatic"}
		types["3"] = Data{Trim: "Si", Transmission: "Automatic"}
		types["4"] = Data{Trim: "Si", Extra: "w/optional 4WS", Transmission: "Automatic"}
	case "CRX", "CR-X":

		if year == 1989 {
			types["5"] = Data{Trim: "STD", Transmission: "Automatic"}
		} else {
			types["5"] = Data{Trim: "STD", Transmission: "Manual"}
		}

		// No idea how to differentiate "HF Civic CRX Manual Seat Belt | Si Civic CRX Manual Seat Belt | Si Civic CRX Automatic ('passive') Seat Belt |
		types["6"] = Data{Trim: "HF, Si"}
	case "Civic":

		if contains(dirt.BodyStyles, "Hatchback") {
			types["4"] = Data{Trim: "STD", Transmission: "Manual"}
			types["5"] = Data{Trim: "DX", Transmission: "Manual"}
			types["6"] = Data{Trim: "Si", Transmission: "Manual"}
		} else if contains(dirt.BodyStyles, "Sedan") || contains(dirt.Doors, "4 Door") {
			types["4"] = Data{Trim: "DX", Transmission: "Manual"}
			types["5"] = Data{Trim: "LX", Transmission: "Manual"}
		} else if contains(dirt.BodyStyles, "Wagon") || contains(dirt.BodyStyles, "Wagovan") {
			types["4"] = Data{Trim: "STD", Transmission: "Manual"}
			types["5"] = Data{Trim: "STD", Transmission: "Manual"}
			types["6"] = Data{Trim: "RTi", Extra: "4WD", Transmission: "Manual"}
		}
	}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func gradeTypePre93(typeCode, model string, year int, dirt *Dirty) {
	types := make(map[string]Data)

	switch model {
	case "Accord":
		types["4"] = Data{Trim: "DX"}
		types["5"] = Data{Trim: "LX"}
		types["6"] = Data{Trim: "DC"}

		if year >= 1992 && year <= 1993 {
			types["7"] = Data{Trim: "EX"}
		}

		types["8"] = Data{Trim: "SE"}
		types["9"] = Data{Trim: "10th Anniversary"}
	case "Prelude":
		types["1"] = Data{Trim: "S", EngineSize: "2.0L"}
		types["2"] = Data{Trim: "S", EngineSize: "2.0L"}
		types["3"] = Data{Trim: "S", EngineSize: "2.1L"}

		if year >= 1992 && year <= 1993 {
			types["4"] = Data{Trim: "S"}
			types["6"] = Data{Trim: "Si", Extra: "w/4WS"}
			types["7"] = Data{Extra: "VTEC"}
		} else {
			types["4"] = Data{Trim: "Si", Extra: "w/4WS", EngineSize: "2.1L"}
		}

		types["5"] = Data{Trim: "Si", EngineSize: "2.3L"}
	case "del Sol":
		if year == 1993 {
			types["4"] = Data{Trim: "S"}
			types["6"] = Data{Trim: "Si"}
		}
	case "Civic":
		if year == 1992 {
			types["4"] = Data{Trim: "CX", Doors: "3 Door"}
			types["6"] = Data{Trim: "VX"}
		} else {
			types["4"] = Data{Trim: "STD, DX", Doors: "3 Door"}
		}

		if year == 1993 {
			types["5"] = Data{Trim: "CX"}

			if contains(dirt.Doors, "3 Door") {
				types["6"] = Data{Trim: "DX"}
			} else if contains(dirt.Doors, "2 Door") {
				types["6"] = Data{Trim: "EX"}
			}

			types["7"] = Data{Trim: "VX"}
		} else {
			if contains(dirt.Doors, "3 Door") {
				types["5"] = Data{Trim: "DX"}
				types["6"] = Data{Trim: "Si"}
			} else if contains(dirt.BodyStyles, "Wagon") {
				types["5"] = Data{Trim: "DX", Extra: "2WD"}
				types["6"] = Data{Trim: "ATi", Extra: "4WD"}
			} else {
				types["5"] = Data{Trim: "LX"}
				types["6"] = Data{Trim: "EX"}
			}
		}

		types["8"] = Data{Trim: "Si", Doors: "3 Door"}
	case "CRX":
		types["6"] = Data{Trim: "HF, Si"}
	}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func gradeTypePre99(typeCode, model string, year int, dirt *Dirty) {
	types := make(map[string]Data)

	switch model {
	case "Accord":
		types["0"] = Data{Trim: "SE", Doors: "2 Door, 4 Door"}
		types["1"] = Data{Trim: "DX", Extra: "w/ABS", Doors: "2 Door, 4 Door"}

		if contains(dirt.BodyStyles, "Wagon") || contains(dirt.BodyStyles, "Wagovan") {
			types["2"] = Data{Trim: "LX"}
			types["3"] = Data{Trim: "LX", Extra: "w/ABS"}
			types["9"] = Data{Trim: "EX"}
		} else {
			types["2"] = Data{Trim: "DX", Doors: "2 Door, 4 Door"}
			types["3"] = Data{Trim: "LX", Extra: "w/ABS, V6", Doors: "2 Door, 4 Door"}
			types["9"] = Data{Trim: "Value Package", Doors: "4 Door"}
		}

		if year == 1998 {
			types["4"] = Data{Trim: "DX", Doors: "4 Door"}
			types["7"] = Data{Trim: "EX", Extra: "ULEV"}
		} else {
			types["4"] = Data{Trim: "LX", Extra: "w/ABS", Doors: "2 Door, 4 Door"}
			types["7"] = Data{Trim: "EX"}
		}

		types["5"] = Data{Trim: "EX", Doors: "2 Door, 4 Door"}
		types["6"] = Data{Trim: "E", Extra: "w/Leather", Doors: "2 Door, 4 Door"}
		types["8"] = Data{Trim: "DX 25th Anniversary"}
	case "Prelude":
		if year == 1997 {
			types["4"] = Data{Extra: "VTEC"}
			types["5"] = Data{Trim: "SH"}
		} else if year == 1998 {
			types["4"] = Data{Extra: "w/o SH pkg."}
			types["5"] = Data{Trim: "Si"}
		} else {
			types["4"] = Data{Trim: "S"}
			types["5"] = Data{Trim: "Si"}
		}

		types["6"] = Data{Trim: "Si", Extra: "4WS"}
		types["7"] = Data{Extra: "VTEC"}
	case "del Sol":
		types["4"] = Data{Trim: "S"}
		types["6"] = Data{Trim: "Si"}
		types["7"] = Data{Trim: "Si", Extra: "w/ABS, VTEC"}
		types["9"] = Data{Trim: "LX, Si", Extra: "VTEC, w/ABS", Doors: "3 Door"}
	case "Civic":
		types["0"] = Data{Trim: "LX"}
		types["2"] = Data{Trim: "CX, DX, EX, HX, CVT", Doors: "2 Door, 3 Door, 4 Door"}
		types["3"] = Data{Trim: "CX, DX, EX, LX", Extra: "w/AC, w/ABS", Doors: "2 Door, 3 Door, 4 Door"}
		types["4"] = Data{Trim: "DX, EX", Doors: "2 Door, 4 Door"}
		types["5"] = Data{Trim: "CX, LX, DX, EX", Extra: "w/AC, w/ABS", Doors: "2 Door, 3 Door, 4 Door"}
		types["6"] = Data{Trim: "DX, LX", Extra: "w/ABS", Doors: "3 Door, 4 Door"}
		types["7"] = Data{Trim: "LX, VX", Extra: "w/AC", Doors: "3 Door"}
		types["8"] = Data{Trim: "Si, LX", Extra: "w/ABS", Doors: "3 Door"}
		types["9"] = Data{Trim: "Si, EX", Extra: "w/ABS", Doors: "3 Door, 4 Door"}
	case "CR-V":
		types["4"] = Data{Trim: "LX", Extra: "w/o ABS"}
		types["5"] = Data{Extra: "w/ABS"}

		if year >= 1998 && year <= 1999 {
			types["6"] = Data{Trim: "EX"}
		}
	case "Odyssey":
		types["4"] = Data{Trim: "LX", Extra: "6 Seater Passaway"}
		types["6"] = Data{Trim: "LX", Extra: "7 Seater Passenger"}
		types["7"] = Data{Trim: "EX"}
	}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func findHondaAssemblyPlant(typeCode string) string {
	types := make(map[string]string)

	types["A"] = "Marysville, Ohio, USA"
	types["B"] = "Lincoln, Alabama, USA"
	types["C"] = "Sayama, Saitama, Japan"
	types["E"] = "Greensburg, Indiana, USA"
	types["G"] = "El Salto, Jalisco, Mexico"
	types["H"] = "Alliston, Ontario, Canada"
	types["L"] = "East Liberty, Ohio, USA"
	types["M"] = "Celaya, Guanajuato, Mexico"
	types["P"] = "Ayutthaya, Thailand"
	types["S"] = "Suzuka, Mie, Japan"
	types["T"] = "Utsunomiya, Tochigi, Japan"
	types["U"] = "Swindon, Wiltshire, U.K."

	result := "Unknown"

	if val, ok := types[typeCode]; ok {
		result = val
	}

	return result
}

func cleanModel(rawModel []string) string {
	result := "Unknown"

	models := []string{
		"Accord",
		"Civic",
		"Prelude",
		"CR-V",
		"CRX",
		"CR-X",
		"Odyssey",
		"Orthia",
	}

	for _, v := range models {
		if contains(rawModel, v) {
			result = v
			break
		}
	}

	return result
}

/*
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
}*/
