package control

import "testing"

func TestApplication_GetPagedData_Page3(t *testing.T) {
	expect := 3
	actual, _ := getPageData("C5")

	if actual != expect {
		t.Errorf("Expecting %v, got %v", expect, actual)
	}
}

func TestApplication_GetPagedData_Size5(t *testing.T) {
	expect := 5
	_, actual := getPageData("A5")

	if actual != expect {
		t.Errorf("Expecting %v, got %v", expect, actual)
	}
}

func TestApplication_GetPagedData_Page26(t *testing.T) {
	expect := 26
	actual, _ := getPageData("Z5")

	if actual != expect {
		t.Errorf("Expecting %v, got %v", expect, actual)
	}
}
