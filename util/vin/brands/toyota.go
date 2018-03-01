package brands

import "github.com/louisevanderlith/mango/util/vin/common"

type Toyota struct {
	common.VDS
}

type YearHold struct {
	Start int
	End   int
	Data  Data
}

// http://vehicleidentificationnumber.com/vin-decoder/toyota-car.html
func (v Toyota) GetPassengerCar(sections common.VINSections, year int) common.VDS {
	return v.VDS
}

// pos 4
/*func findEngineType(typeCode string, year int, dirt *Dirty) {
	types := make(map[string][]YearHold)

	types["A"] = append(types["A"], YearData(1983, 1988, Data{EngineSize: "1.5", EngineModel: "3A-C", Extra: "I4"}))
	types["A"] = append(types["A"], YearData(1983, 1987, Data{EngineSize: "1.6", EngineModel: "4A-C", Extra: "I4"}))
	types["A"] = append(types["A"], YearData(1988, 1990, Data{EngineSize: "1.6", EngineModel: "4A-F", Extra: "I4, Carb"}))
	types["A"] = append(types["A"], YearData(1989, 1995, Data{EngineSize: "1.6", EngineModel: "4A-FE", Extra: "I4"}))
	types["A"] = append(types["A"], YearData(1988, 1988, Data{EngineSize: "1.6", EngineModel: "4A-GE", Extra: "I4"}))
	types["A"] = append(types["A"], YearData(1985, 1987, Data{EngineSize: "1.6", EngineModel: "4A-GEC", Extra: "I4"}))
	types["A"] = append(types["A"], YearData(1987, 1987, Data{EngineSize: "1.6", EngineModel: "4A-GELC", Extra: "I4, Turbo"}))
	types["A"] = append(types["A"], YearData(1985, 1987, Data{EngineSize: "1.6", EngineModel: "4A-LC", Extra: "I4, Turbo"}))
	types["A"] = append(types["A"], YearData(1988, 1989, Data{EngineSize: "1.6", EngineModel: "4A-GZE", Extra: "I4"}))
	types["A"] = append(types["A"], YearData(1993, 1995, Data{EngineSize: "1.8", EngineModel: "7A-FE", Extra: "I4"}))

	types["C"] = append(types["C"], YearData(1984, 1985, Data{EngineSize: "1.8", EngineModel: "1C-LC", Extra: "I4, Diesel"}))
	types["C"] = append(types["C"], YearData(1984, 1986, Data{EngineSize: "1.8", EngineModel: "1C-TLC", Extra: "I4, Turbo"}))

	types["E"] = append(types["E"], YearData(1987, 1990, Data{EngineSize: "1.5", EngineModel: "3-E", Extra: "I4, Carb, Turbo"}))
	types["E"] = append(types["E"], YearData(1990, 1994, Data{EngineSize: "1.5", EngineModel: "3-EE", Extra: "I4, EFI"}))
	types["E"] = append(types["E"], YearData(1992, 1995, Data{EngineSize: "1.5", EngineModel: "5E-FE", Extra: "I4"}))

	types["K"] = append(types["K"], YearData(1983, 1983, Data{EngineSize: "1.3", EngineModel: "4K-C", Extra: "I4, Carb"}))
	types["K"] = append(types["K"], YearData(1983, 1984, Data{EngineSize: "1.3", EngineModel: "4K-E", Extra: "I4, EFI"}))

	types["M"] = append(types["M"], YearData(1983, 1988, Data{EngineSize: "2.8", EngineModel: "5M-GE", Extra: "I6"}))
	types["M"] = append(types["M"], YearData(1987, 1992, Data{EngineSize: "3.0", EngineModel: "7M-GE", Extra: "I6"}))
	types["M"] = append(types["M"], YearData(1987, 1992, Data{EngineSize: "3.0", EngineModel: "7M-GET", Extra: "I6, Turbo"}))

	types["R"] = append(types["R"], YearData(1983, 1984, Data{EngineSize: "2.4", EngineModel: "22-RE", Extra: "I4"}))

	types["S"] = append(types["S"], YearData(1983, 1986, Data{EngineSize: "2.0", EngineModel: "2S-ELC", Extra: "I4"}))
	types["S"] = append(types["S"], YearData(1987, 1991, Data{EngineSize: "2.0", EngineModel: "3S-FE", Extra: "I4"}))
	types["S"] = append(types["S"], YearData(1988, 1989, Data{EngineSize: "2.0", EngineModel: "3S-GE", Extra: "I4"}))
	types["S"] = append(types["S"], YearData(1986, 1987, Data{EngineSize: "2.0", EngineModel: "3S-GEC", Extra: "I4"}))
	types["S"] = append(types["S"], YearData(1987, 1995, Data{EngineSize: "2.0", EngineModel: "3S-GELC", Extra: "I4, Turbo"}))
	types["S"] = append(types["S"], YearData(1988, 1990, Data{EngineSize: "2.0", EngineModel: "3S-GTE", Extra: "I4"}))
	types["S"] = append(types["S"], YearData(1990, 1995, Data{EngineSize: "2.2", EngineModel: "5S-FE", Extra: "I4"}))

	types["V"] = append(types["V"], YearData(1983, 1988, Data{EngineSize: "2.5", EngineModel: "2VZ-FE", Extra: "V6"}))
	types["V"] = append(types["V"], YearData(1983, 1988, Data{EngineSize: "3.0", EngineModel: "3VZ-FE", Extra: "V6"}))

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}*/

// pos 5
/*func findVehicleLine(typeCode string, year int, dirt *Dirty) {
	types := make(map[string][]YearHold)

	types["A"] = append(types["A"], YearData(1983, 1985, Data{Model: "Celica"}))
	types["A"] = append(types["A"], YearData(1984, 1995, Data{Model: "Supra"}))
	types["B"] = append(types["B"], YearData(1995, 1999, Data{Model: "Avalon"}))
	types["E"] = append(types["E"], YearData(1983, 1985, Data{Model: "Corolla"}))
	types["K"] = append(types["K"], YearData(1992, 1995, Data{Model: "Camry"}))
	types["L"] = append(types["L"], YearData(1983, 1995, Data{Model: "Tercel"}))
	types["L"] = append(types["L"], YearData(1992, 1995, Data{Model: "Paseo"}))
	types["P"] = append(types["P"], YearData(1983, 1984, Data{Model: "Starlet"}))
	types["T"] = append(types["T"], YearData(1986, 1995, Data{Model: "Celica"}))
	types["V"] = append(types["V"], YearData(1983, 1991, Data{Model: "Camry"}))
	types["W"] = append(types["W"], YearData(1985, 1989, Data{Model: "MR2"}))
	types["W"] = append(types["W"], YearData(1991, 1995, Data{Model: "MR2"}))
	types["X"] = append(types["X"], YearData(1983, 1990, Data{Model: "Cressida"}))

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}*/

func YearData(startYear, endYear int, data Data) YearHold {
	var result YearHold

	result = YearHold{
		Start: startYear,
		End:   endYear,
		Data:  data,
	}

	return result
}

/*
func groupsToyota() {
	const toyota = "Toyota"
	const hino = "Hino"
	const daihatsu = "Daihatsu"
	const lexus = "Lexus"
	descrip := Toyota{}

	groupj := NewWMIGroup("J")
	groupj.Add("TD", toyota, PassengerCar, descrip)
	groupj.Add("TG", toyota, Bus, descrip)
	groupj.Add("TK", toyota, PassengerCar, descrip)
	groupj.Add("TN", toyota, PassengerCar, descrip)
	groupj.Add("TX", toyota, PassengerCar, descrip)
	groupj.Add("T1", toyota, PassengerCar, descrip)
	groupj.Add("T2", toyota, PassengerCar, descrip)
	groupj.Add("T7", toyota, PassengerCar, descrip)
	groupj.Add("TA", toyota, Truck, descrip)
	groupj.Add("TB", toyota, Truck, descrip)
	groupj.Add("TF", toyota, Truck, descrip)
	groupj.Add("TM", toyota, Truck, descrip)
	groupj.Add("T4", toyota, Truck, descrip)
	groupj.Add("TE", toyota, MPV, descrip)
	groupj.Add("TL", toyota, MPV, descrip)
	groupj.Add("T3", toyota, MPV, descrip)
	groupj.Add("T5", toyota, IncompleteCar, descrip)
	groupj.Add("TH", lexus, PassengerCar, descrip)
	groupj.Add("TJ", lexus, MPV, descrip)
	groupj.Add("T6", lexus, MPV, descrip)
	groupj.Add("T8", lexus, PassengerCar, descrip)
	groupj.Add("DA", daihatsu, PassengerCar, descrip)
	groupj.Add("D1", daihatsu, PassengerCar, descrip)
	groupj.Add("D2", daihatsu, PassengerCar, descrip)
	groupj.Add("HA", hino, Truck, descrip)
	groupj.Add("HB", hino, Truck, descrip)
	groupj.Add("HC", hino, Truck, descrip)
	groupj.Add("HD", hino, Truck, descrip)
	groupj.Add("HE", hino, Truck, descrip)
}
*/
