package common

import "testing"

func TestIsValid_Correct(t *testing.T) {
	input := "1G1BL52P7TR115520"
	checkInput := "7"
	actual := IsValid(input, checkInput)

	if !actual {
		t.Errorf("VIN is not valid.")
	}
}
