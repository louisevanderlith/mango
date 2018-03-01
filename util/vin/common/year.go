package common

import (
	"strconv"
)

const baseYear int = 1979

// GetYear will return the model year using the 7th character to identify generation.
func GetYear(yearCode, char7 string) int {
	var result int

	years := getAllYears(yearCode)
	isPre2009 := isYearPre2009(char7)

	if isPre2009 {
		result = years[0]
	} else {
		result = years[1]
	}

	return result
}

func isYearPre2009(char7 string) bool {
	// Char7 is Numeric, then year will be pre 2009
	_, err := strconv.Atoi(char7)
	return err == nil
}

func getAllYears(code string) []int {
	value := getValue(code, "I", "O", "Q", "U", "Z", "0")
	lowYear := baseYear + value
	highYear := lowYear + 30

	return []int{lowYear, highYear}
}
