package common

import "testing"

func TestGetClassByWeight_ClassB(t *testing.T) {
	expect := "B"
	actual, err := getClassByWeight(1813)

	if err != nil || actual.Name != expect {
		t.Errorf("Expecting Class %s, found Class %s. Error: %s", expect, actual.Name, err)
	}
}

func TestGetClassByWeight_InvalidWeight(t *testing.T) {
	_, err := getClassByWeight(-100)

	if err == nil {
		t.Error("Invalid weight didn't return Error.")
	}
}
