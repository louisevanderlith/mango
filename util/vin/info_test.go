package vin

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/louisevanderlith/mango/util/vin/common"
)

func TestLoadSectionCodes_Correct(t *testing.T) {
	input := "JHLRD68405C200888"
	expected := common.VINSections{
		WMICode:           "JHL",
		FeatureCode:       "RD684",
		CheckDigit:        "0",
		YearCode:          "5",
		AssemblyPlantCode: "C",
		SequenceCode:      "200888",
		FullVIN:           input,
	}

	actual, _ := common.LoadVINSections(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetInfo_Correct(t *testing.T) {
	input := "JHLRD68405C200888"
	actual, _ := GetInfo(input)

	fmt.Println("ACT", actual)

	t.Fail()
}

func TestLoadWMI(t *testing.T) {
	input, err := common.LoadVINSections("JHLRD68405C200888")

	if err != nil {
		t.Error(err)
	}

	expected := common.WMI{
		Manufacturer: common.Manufacturer{
			Category: common.NotSpecified,
			Name:     "Honda",
			VDSName:  "Honda",
		},
		Region: common.Region{
			Continent: "Asia",
			Country:   "Japan",
		},
	}

	actual := loadWMI(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestLoadVDS(t *testing.T) {
	input, err := common.LoadVINSections("1HGEJ8144XL019972")

	if err != nil {
		t.Error(err)
	}

	expected := common.VDS{
		Model:         "Civic",
		BodyStyle:     "Coupe",
		Doors:         "2 Door",
		EngineModel:   "",
		EngineSize:    "1.6L",
		Extras:        []string{},
		Gears:         "5",
		Transmission:  "Manual",
		Trim:          "EX",
		AssemblyPlant: "East Liberty, Ohio, USA",
	}

	actual := loadVDS(input, "honda", 1999)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestLoadVIS(t *testing.T) {
	input, err := common.LoadVINSections("1HGEJ8144XL019972")

	if err != nil {
		t.Error(err)
	}

	expected := common.VIS{
		SequenceNo: "019972",
		ValidVIN:   true,
		Year:       1999,
	}

	actual := loadVIS(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
