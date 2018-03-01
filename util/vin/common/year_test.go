package common

import (
	"reflect"
	"testing"
)

func TestGetYear_A_19802010(t *testing.T) {
	input := "A"
	expected := []int{1980, 2010}
	actual := GetYear(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expecting %v got %v", expected, actual)
	}
}

func TestGetYear_G_19862016(t *testing.T) {
	input := "G"
	expected := []int{1986, 2016}
	actual := GetYear(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expecting %v got %v", expected, actual)
	}
}

func TestGetYear_3_20032033(t *testing.T) {
	input := "3"
	expected := []int{2003, 2033}
	actual := GetYear(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expecting %v got %v", expected, actual)
	}
}
