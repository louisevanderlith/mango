package vin

import (
	"testing"

	"github.com/louisevanderlith/mango/util/vin/brands"
	"github.com/louisevanderlith/mango/util/vin/common"
)

func TestGetData_Honda_Civic_Success(t *testing.T) {
	input, _ := common.LoadVINSections("1HGEJ8144XL019972")
	year := common.GetYear(input.VISCode.YearCode, input.VDSCode.Char7)
	actual := brands.Honda{}.GetPassengerCar(input, year)

	t.Log(actual)
	t.Fail()
}

func TestGetData_Honda_CRV_Success(t *testing.T) {
	input, _ := common.LoadVINSections("JHLRD68405C200888")
	year := common.GetYear(input.VISCode.YearCode, input.VDSCode.Char7)
	actual := brands.Honda{}.GetPassengerCar(input, year)

	t.Log(actual)
	t.Fail()
}

func TestGetData_Honda_Acura_Success(t *testing.T) {
	input, _ := common.LoadVINSections("JH4NA1158NT000999")
	year := common.GetYear(input.VISCode.YearCode, input.VDSCode.Char7)
	actual := brands.Honda{}.GetPassengerCar(input, year)

	t.Log(actual)
	t.Fail()
}

func TestGetData_Volvo_Success(t *testing.T) {
	input, _ := common.LoadVINSections("YV1RS592962540277")
	year := common.GetYear(input.VISCode.YearCode, input.VDSCode.Char7)
	actual := brands.Volvo{}.GetPassengerCar(input, year)

	t.Log(actual)
	t.Fail()
}

func TestGetData_VW_Success(t *testing.T) {
	input, _ := common.LoadVINSections("WVWZZZ3CZEE140287")
	year := common.GetYear(input.VISCode.YearCode, input.VDSCode.Char7)
	actual := brands.Volkswagen{}.GetPassengerCar(input, year)

	t.Log(actual)
	t.Fail()
}

/* VW & Volvo VIN
'WVWZZZ3CZEE140287' - VW,
'U5YFF52229L068332' - KIA,
'1M8GDM9AXKP042788' - Motor Coach Industries??,
'KMFZBX7HLCU835815' - Hyundai Truck,
'5N1AR2MM9GC610938' - Nissan,
'U5YHM516ADL012319' - KIA,
'WAUZZZ8R0BA098735' - Audi,
'VF3WC5FS9AW055905' - Peugeot,
'WDDEJ71X18A016288' - Mercedes,
'VF7RERHRH76843233' - Citroen,
'WDB2110161A032195' - Mercedes,
'VF77ENFUXFJ518177' - Citroen,
'W0L0ZCF6848022377' - Opel,
'WDDEJ76X28A010315' - Mercedes
*/
