package common

import (
	"testing"
)

const fullVIN string = "1HGEJ8144XL019972"

var sections VINSections

func init() {
	sections, _ = LoadVINSections(fullVIN)
}

func TestWMICode_Full(t *testing.T) {
	expect := "1HG"
	actual := sections.WMICode.FullWMI

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestWMICode_ContinentCode(t *testing.T) {
	expect := "1"
	actual := sections.WMICode.ContinentCode

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestWMICode_RegionCode(t *testing.T) {
	expect := "H"
	actual := sections.WMICode.RegionCode

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestWMICode_ManufacturerCode(t *testing.T) {
	expect := "HG"
	actual := sections.WMICode.ManufacturerCode

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestVDSCode_Full(t *testing.T) {
	expect := "EJ814"
	actual := sections.VDSCode.FullVDS

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestVDSCode_Char4(t *testing.T) {
	expect := "E"
	actual := sections.VDSCode.Char4

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestVDSCode_Char5(t *testing.T) {
	expect := "J"
	actual := sections.VDSCode.Char5

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestVDSCode_Char6(t *testing.T) {
	expect := "8"
	actual := sections.VDSCode.Char6

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestVDSCode_Char7(t *testing.T) {
	expect := "1"
	actual := sections.VDSCode.Char7

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestVDSCode_Char8(t *testing.T) {
	expect := "4"
	actual := sections.VDSCode.Char8

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestVDSCode_GetTypeCode(t *testing.T) {
	expect := "J81"
	actual := sections.VDSCode.GetTypeCode("5", "6", "7")

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestVISCode_Full(t *testing.T) {
	expect := "4XL019972"
	actual := sections.VISCode.FullVIS

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestVISCode_CheckDigit(t *testing.T) {
	expect := "4"
	actual := sections.VISCode.CheckDigit

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestVISCode_YearCode(t *testing.T) {
	expect := "X"
	actual := sections.VISCode.YearCode

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestVISCode_AssemblyCode(t *testing.T) {
	expect := "L"
	actual := sections.VISCode.AssemblyPlantCode

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}

func TestVISCode_SeqCode(t *testing.T) {
	expect := "019972"
	actual := sections.VISCode.SequenceCode

	if expect != actual {
		t.Errorf("Expected %s, got %s", expect, actual)
	}
}
