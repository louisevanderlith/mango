package common

import (
	"time"
)

const baseYear int = 1979

// GetYear will return the year(s) relating to the code character.
// This function will always return 2 year options.
// A = 1980 & 2010
func GetYear(code string) []int {
	value := getValue(code, "I", "O", "Q", "U", "Z", "0")
	lowYear := baseYear + value
	highYear := lowYear + 30

	return []int{lowYear, highYear}
}

// GetBGYear (Best Guess Year) will return the year which is most likely to be correct.
func GetBGYear(code string) int {
	var result int

	years := GetYear(code)
	now := time.Now().Year()

	for _, v := range years {
		if v <= now {
			result = v
			break
		}
	}

	return result
}
