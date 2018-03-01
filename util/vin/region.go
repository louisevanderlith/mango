package vin

import (
	"strings"
)

type Region struct {
	Continent string
	Country   string
}

type Continent struct {
	CharKey
	Countries []Country
}

type Country struct {
	CharKey
	Prefix string
}

type CharKey struct {
	Name  string
	Start string
	End   string
}

var world []Continent

func init() {
	world = buildWorld()
}

func getRegion(wmi string) Region {
	var result Region

	if len(wmi) == 2 {
		const illegalChars = "IOQ"

		continentCode := wmi[:1]
		continentValid := strings.Index(illegalChars, continentCode) == -1
		countryCode := wmi[1:]
		countryValid := strings.Index(illegalChars, countryCode) == -1

		if continentValid && countryValid {
			continent := findContinent(continentCode)

			if len(continent.Countries) > 0 {
				county := findCountry(continent.Countries, countryCode)

				result.Continent = continent.Name
				result.Country = county.Name
			}
		}
	}

	return result
}

func findContinent(continentCode string) Continent {
	var result Continent

	for _, v := range world {
		if isBetween(continentCode, v.Start, v.End) {
			result = v
			break
		}
	}

	return result
}

func findCountry(countries []Country, countryCode string) Country {
	var result Country

	for _, v := range countries {
		if isBetween(countryCode, v.Start, v.End) {
			result = v
			break
		}
	}

	return result
}

func isBetween(char, start, end string) bool {
	// Check to see if a character is between to other characters in sequence.
	charWeight := getWeight(char)
	startWeight := getWeight(start)
	endWeight := getWeight(end)

	return startWeight <= charWeight && endWeight >= charWeight
}

func buildWorld() []Continent {
	var result []Continent

	africa := buildContinent("Africa", "A", "H", getAfricaRegions())
	asia := buildContinent("Asia", "J", "R", getAsiaRegions())
	europe := buildContinent("Europe", "S", "Z", getEuropeRegions())
	na := buildContinent("North America", "1", "5", getNARegions())
	oceania := buildContinent("Oceania", "6", "7", getOceaniaRegions())
	sa := buildContinent("South America", "8", "0", getSARegions())

	result = append(result, africa, asia, europe, na, oceania, sa)

	return result
}

func getAfricaRegions() []Country {
	var result []Country

	sa := buildCountry("A", "South Africa", "A", "H")
	ic := buildCountry("A", "Ivory Coast", "J", "N")
	aOther := buildCountry("A", "Other", "P", "0")

	result = append(result, sa, ic, aOther)

	ang := buildCountry("B", "Angola", "A", "E")
	ken := buildCountry("B", "Kenya", "F", "K")
	tan := buildCountry("B", "Tanzania", "L", "R")
	bOther := buildCountry("B", "Other", "S", "0")
	result = append(result, ang, ken, tan, bOther)

	ben := buildCountry("C", "Benin", "A", "E")
	mad := buildCountry("C", "Madagascar", "F", "K")
	tun := buildCountry("C", "Tunisia", "L", "R")
	cOther := buildCountry("C", "Other", "S", "0")
	result = append(result, ben, mad, tun, cOther)

	egp := buildCountry("D", "Egypt", "A", "E")
	mor := buildCountry("D", "Morocco", "F", "K")
	zam := buildCountry("D", "Zambia", "L", "R")
	dOther := buildCountry("D", "Other", "S", "0")
	result = append(result, egp, mor, zam, dOther)

	gha := buildCountry("E", "Ghana", "A", "E")
	moz := buildCountry("E", "Mozambique", "F", "K")
	eOther := buildCountry("E", "Other", "L", "0")
	result = append(result, gha, moz, eOther)

	ind := buildCountry("F", "India", "A", "E")
	nig := buildCountry("F", "Nigeria", "F", "K")
	fOther := buildCountry("F", "Other", "L", "0")
	result = append(result, ind, nig, fOther)

	gOther := buildCountry("G", "Other", "A", "0")
	hOther := buildCountry("H", "Other", "A", "0")
	result = append(result, gOther, hOther)

	return result
}

func getAsiaRegions() []Country {
	var result []Country

	jpa := buildCountry("J", "Japan", "A", "0")
	result = append(result, jpa)

	sri := buildCountry("K", "Sri Lanka", "A", "E")
	isr := buildCountry("K", "Isreal", "F", "K")
	sko := buildCountry("K", "South Korea", "L", "R")
	kza := buildCountry("K", "Kazakhstan", "S", "0")
	result = append(result, sri, isr, sko, kza)

	chi := buildCountry("L", "China", "A", "0")
	result = append(result, chi)

	ind := buildCountry("M", "India", "A", "E")
	indo := buildCountry("M", "Indonesia", "F", "K")
	tha := buildCountry("M", "Thailand", "L", "R")
	mya := buildCountry("M", "Myanmar", "S", "S")
	mOther := buildCountry("M", "Other", "T", "0")
	result = append(result, ind, indo, tha, mya, mOther)

	ira := buildCountry("N", "Iran", "A", "E")
	pak := buildCountry("N", "Pakistan", "F", "K")
	tur := buildCountry("N", "Turkey", "L", "R")
	nOther := buildCountry("N", "Other", "S", "0")
	result = append(result, ira, pak, tur, nOther)

	phi := buildCountry("P", "Philippines", "A", "E")
	sin := buildCountry("P", "Singapore", "F", "K")
	mal := buildCountry("P", "Malaysia", "L", "R")
	pOther := buildCountry("P", "Other", "S", "0")
	result = append(result, phi, sin, mal, pOther)

	uae := buildCountry("R", "UAE", "A", "E")
	tai := buildCountry("R", "Taiwan", "F", "K")
	vie := buildCountry("R", "Vietnam", "L", "R")
	sau := buildCountry("R", "Saudi Arabia", "S", "0")
	result = append(result, uae, tai, vie, sau)

	return result
}

func getEuropeRegions() []Country {
	var result []Country

	uk := buildCountry("S", "United Kingdom", "A", "M")
	eger := buildCountry("S", "Eastern Germany", "N", "T")
	pol := buildCountry("S", "Poland", "U", "Z")
	lat := buildCountry("S", "Latvia", "1", "4")
	sOther := buildCountry("S", "Other", "5", "0")
	result = append(result, uk, eger, pol, lat, sOther)

	swi := buildCountry("T", "Switzerland", "A", "H")
	cze := buildCountry("T", "Czech Republic", "J", "P")
	hun := buildCountry("T", "Hungary", "R", "V")
	por := buildCountry("T", "Portugal", "W", "1")
	tOther := buildCountry("T", "Other", "2", "0")
	result = append(result, swi, cze, hun, por, tOther)

	uOther := buildCountry("U", "Other", "A", "G")
	den := buildCountry("U", "Denmark", "H", "M")
	ire := buildCountry("U", "Ireland", "N", "T")
	rom := buildCountry("U", "Romania", "U", "Z")
	uOtherB := buildCountry("U", "Other", "1", "4")
	slo := buildCountry("U", "Slovakia", "5", "7")
	uOtherC := buildCountry("U", "Other", "8", "0")
	result = append(result, uOther, den, ire, rom, uOtherB, slo, uOtherC)

	aus := buildCountry("V", "Austria", "A", "E")
	fra := buildCountry("V", "France", "F", "R")
	spa := buildCountry("V", "Spain", "S", "W")
	ser := buildCountry("V", "Serbia", "X", "2")
	cro := buildCountry("V", "Croatia", "3", "5")
	est := buildCountry("V", "Estonia", "6", "0")
	result = append(result, aus, fra, spa, ser, cro, est)

	ger := buildCountry("W", "Germany", "A", "0")
	result = append(result, ger)

	bul := buildCountry("X", "Bulgaria", "A", "E")
	gre := buildCountry("X", "Greece", "F", "K")
	net := buildCountry("X", "Netherlands", "L", "R")
	ussr := buildCountry("X", "USSR", "S", "W")
	lux := buildCountry("X", "Luxembourg", "X", "2")
	rus := buildCountry("X", "Russia", "3", "0")
	result = append(result, bul, gre, net, ussr, lux, rus)

	bel := buildCountry("Y", "Belgium", "A", "E")
	fin := buildCountry("Y", "Finland", "F", "K")
	mal := buildCountry("Y", "Malta", "L", "R")
	swe := buildCountry("Y", "Sweden", "S", "W")
	nor := buildCountry("Y", "Norway", "X", "2")
	bela := buildCountry("Y", "Belarus", "3", "5")
	ukr := buildCountry("Y", "Ukraine", "6", "0")
	result = append(result, bel, fin, mal, swe, nor, bela, ukr)

	ita := buildCountry("Z", "Italy", "A", "R")
	zOther := buildCountry("Z", "Other", "S", "W")
	slov := buildCountry("Z", "Slovenia", "X", "2")
	lit := buildCountry("Z", "Lithuania", "3", "5")
	zOtherB := buildCountry("Z", "Other", "6", "0")
	result = append(result, ita, zOther, slov, lit, zOtherB)

	return result
}

func getNARegions() []Country {
	var result []Country

	usa := buildCountry("1", "United States", "A", "0")
	can := buildCountry("2", "Canada", "A", "0")
	mex := buildCountry("3", "Mexico", "A", "0")
	usab := buildCountry("4", "United States", "A", "0")
	usac := buildCountry("5", "United States", "A", "0")

	result = append(result, usa, can, mex, usab, usac)

	return result
}

func getOceaniaRegions() []Country {
	var result []Country

	aus := buildCountry("6", "Australia", "A", "W")
	other := buildCountry("6", "Other", "X", "0")
	result = append(result, aus, other)

	nze := buildCountry("7", "New Zealand", "A", "E")
	otherB := buildCountry("7", "Other", "F", "0")
	result = append(result, nze, otherB)

	return result
}

func getSARegions() []Country {
	var result []Country

	arg := buildCountry("8", "Argentina", "A", "E")
	chi := buildCountry("8", "Chile", "F", "K")
	ecu := buildCountry("8", "Ecuador", "L", "R")
	per := buildCountry("8", "Peru", "S", "W")
	ven := buildCountry("8", "Venezuela", "X", "2")
	other := buildCountry("8", "Other", "3", "0")
	result = append(result, arg, chi, ecu, per, ven, other)

	bra := buildCountry("9", "Brazil", "A", "E")
	col := buildCountry("9", "Colombia", "F", "K")
	par := buildCountry("9", "Paraguay", "L", "R")
	uru := buildCountry("9", "Uruguay", "S", "W")
	tri := buildCountry("9", "Trinidad & Tobago", "X", "2")
	braB := buildCountry("9", "Brazil", "3", "9")
	otherB := buildCountry("9", "Other", "0", "0")
	result = append(result, bra, col, par, uru, tri, braB, otherB)

	sa := buildCountry("0", "Other", "A", "0")
	result = append(result, sa)

	return result
}

func buildCountry(prefix, name, start, end string) Country {
	charKey := CharKey{
		name,
		start,
		end,
	}

	return Country{charKey, prefix}
}

func buildContinent(name, start, end string, countries []Country) Continent {
	charKey := CharKey{
		name,
		start,
		end,
	}

	return Continent{charKey, countries}
}
