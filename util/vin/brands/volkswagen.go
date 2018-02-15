package brands

import "github.com/louisevanderlith/mango/util/vin/common"

type Volkswagen struct {
	common.VDS
}

func (v Volkswagen) GetPassengerCar(sections common.VINSections, year int) common.VDS {
	modelCode := sections.VDSCode.GetTypeCode("4", "5")

	dirt := Dirty{}

	findVWModel(modelCode, year, &dirt)

	v.VDS = compileDirty(dirt)
	v.AssemblyPlant = findVWAssemblyPlant(sections.VISCode.AssemblyPlantCode, year)

	return v.VDS
}

func findVWModel(typeCode string, year int, dirt *Dirty) {
	types := make(map[string]Data)

	types["11"] = Data{Model: "Beetle (Brazilian Mexican Nigerian)"}
	types["13"] = Data{Model: "Scirocco 3"}
	types["14"] = Data{Model: "Caddy Mk 1 (European Golf 1 pickup)"}
	types["15"] = Data{Model: "Cabriolet (1980 Beetle Golf 1)"}

	if year < 2012 {
		types["16"] = Data{Model: "Jetta 1, Jetta 2"}
	} else {
		types["16"] = Data{Model: "Beetle"}
	}

	types["17"] = Data{Model: "Golf 1"}
	types["18"] = Data{Model: "Iltis"}
	types["19"] = Data{Model: "Golf 2 (early)"}
	types["1C"] = Data{Model: "New Beetle (US market)"}
	types["1E"] = Data{Model: "Golf 3 Cabriolet"}
	types["1F"] = Data{Model: "Eos"}
	types["1G"] = Data{Model: "Golf and Jetta 2 (late)"}
	types["1H"] = Data{Model: "Golf and Vento 3"}
	types["1J"] = Data{Model: "Golf and Bora 4"}
	types["1K"] = Data{Model: "Golf and Jetta 5 6"}
	types["1T"] = Data{Model: "Touran"}
	types["1Y"] = Data{Model: "New Beetle Cabriolet"}
	types["24"] = Data{Model: "T3 Transporter Single/Double Cab Pickup"}
	types["25"] = Data{Model: "T3 Transporter Van Kombi Bus Caravelle"}
	types["28"] = Data{Model: "LT Transporter 1"}
	types["2D"] = Data{Model: "LT Transporter 2"}
	types["2E"] = Data{Model: "Crafter"}
	types["2H"] = Data{Model: "Amarok"}
	types["2K"] = Data{Model: "Caddy Caddy Maxi 3"}
	types["30"] = Data{Model: "Fox (US model ex-Brazil)"}
	types["31"] = Data{Model: "Passat 2"}
	types["32"] = Data{Model: "Santana sedan"}
	types["33"] = Data{Model: "Passat 2 Variant"}
	types["3A"] = Data{Model: "Passat 3 4"}
	types["3B"] = Data{Model: "Passat 5"}

	if year <= 2010 {
		types["3C"] = Data{Model: "Passat 6"}
	} else if year <= 2015 {
		types["3C"] = Data{Model: "Passat 7 Passat CC"}
	} else {
		types["3C"] = Data{Model: "Passat 8 Passat CC"}
	}

	types["3D"] = Data{Model: "Phaeton"}
	types["50"] = Data{Model: "Corrado (early)"}
	types["53"] = Data{Model: "Scirocco 1 and 2"}
	types["5K"] = Data{Model: "Golf and Jetta 6"}
	types["5M"] = Data{Model: "Golf Plus"}
	types["5N"] = Data{Model: "Tiguan"}
	types["5Z"] = Data{Model: "Fox (Europe)"}
	types["60"] = Data{Model: "Corrado (late)"}
	types["6K"] = Data{Model: "Polo Classic Variant 3"}
	types["6N"] = Data{Model: "Polo 3"}
	types["6R"] = Data{Model: "Polo 5"}
	types["6X"] = Data{Model: "Lupo"}
	types["70"] = Data{Model: "T4 Transporter Vans and Pickups"}
	types["74"] = Data{Model: "Taro"}
	types["7H"] = Data{Model: "T5 Transporter"}
	types["7L"] = Data{Model: "Touareg 1"}
	types["7M"] = Data{Model: "Sharan"}
	types["7P"] = Data{Model: "Touareg 2"}
	types["86"] = Data{Model: "Polo and Derby 1 and 2"}
	types["87"] = Data{Model: "Polo Coupe"}
	types["9C"] = Data{Model: "New Beetle"}
	types["9K"] = Data{Model: "Caddy 2 Van (ex-SEAT Ibiza)"}
	types["9N"] = Data{Model: "Polo 4"}
	types["9U"] = Data{Model: "Caddy 2 Pickup (ex-Skoda Felicia)"}
	types["AA"] = Data{Model: "Up!"}

	if val, ok := types[typeCode]; ok {
		fillDirty(dirt, val)
	}
}

func findVWAssemblyPlant(typeCode string, year int) string {
	types := make(map[string]string)

	types["A"] = "Ingolstadt, Germany"
	types["B"] = "Brussels, Belgium"
	types["C"] = "Chattanooga, USA"
	types["D"] = "Bratislava, Slovakia"
	types["E"] = "Emden, Germany"
	types["F"] = "Ipiranga / Resende, Brazil"
	types["G"] = "Graz, Austria"
	types["H"] = "Hanover, Germany"
	types["K"] = "Osnabrück, Germany"
	types["L"] = "Lagos, Nigeria"
	types["M"] = "Puebla, Mexico"
	types["N"] = "Neckarsulm, Germany"
	types["P"] = "Mosel, Germany or Anchieta, Brazil"
	types["R"] = "Martorell, Spain"
	types["S"] = "Salzgitter, Germany"

	if year <= 1994 {
		types["T"] = "Sarajevo, Yugoslavia"
		types["V"] = "Westmoreland, USA"
	} else {
		types["T"] = "Taubaté, Brazil"
		types["V"] = "Palmela, Portugal"
	}

	types["U"] = "Uitenhage, South Africa"
	types["W"] = "Wolfsburg, Germany"
	types["X"] = "Poznan, Poland"
	types["Y"] = "Pamplona, Spain"
	types["1"] = "Győr, Hungary"
	types["2"] = "Anting, China"
	types["3"] = "Changchun, China"
	types["4"] = "Curitiba, Brazil"
	types["6"] = "Düsseldorf, Germany (Mercedes-Benz)"
	types["7"] = "Ludwigsfelde, Germany (Mercedes-Benz)"
	types["8"] = "Dresden, Germany or General Pacheco, Argentina"

	result := "Unknown"

	if val, ok := types[typeCode]; ok {
		result = val
	}

	return result
}

/*
func groupsVW() {
	const vw = "Volkswagen"
	descrip := Volkswagen{}

	groupa := NewWMIGroup("A")
	groupa.Add("AV", vw, PassengerCar, descrip)

	groupl := NewWMIGroup("L")
	groupl.Add("SV", vw, PassengerCar, descrip)

	groupw := NewWMIGroup("W")
	groupw.Add("VW", vw, PassengerCar, descrip)
	groupw.Add("V1", vw, PassengerCar, descrip)
	groupw.Add("V2", vw, PassengerCar, descrip)
	groupw.Add("V3", vw, PassengerCar, descrip)
	groupw.Add("VG", vw, PassengerCar, descrip)

	groupx := NewWMIGroup("X")
	groupx.Add("W8", vw, PassengerCar, descrip)

	group1 := NewWMIGroup("1")
	group1.Add("VW", vw, PassengerCar, descrip)

	group2 := NewWMIGroup("2")
	group2.Add("V4", vw, PassengerCar, descrip)
	group2.Add("V8", vw, PassengerCar, descrip)

	group3 := NewWMIGroup("3")
	group3.Add("VW", vw, PassengerCar, descrip)

	group8 := NewWMIGroup("8")
	group8.Add("AW", vw, PassengerCar, descrip)

	group9 := NewWMIGroup("9")
	group9.Add("3U", vw, PassengerCar, descrip)
	group9.Add("BW", vw, PassengerCar, descrip)
}
*/
