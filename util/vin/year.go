package vin

const baseYear int = 1979

// GetYear will return the year(s) relating to the code character.
// This function will always return 2 year options.
// A = 1980 & 2010
func GetYear(code string) []int {
	weight := getWeight(code, "I", "O", "Q", "U", "Z", "0")
	lowYear := baseYear + weight
	highYear := lowYear + 30

	return []int{lowYear, highYear}
}
