package common

import (
	"strconv"
)

func IsValid(fullVIN, checkChar string) bool {
	sum := 0

	for k, v := range fullVIN {
		weight := getVINWeight(k)
		value := getCharValue(string(v))
		sum += weight * value
	}

	remainder := strconv.Itoa(sum % 11)

	if remainder == "10" {
		remainder = "X"
	}

	return remainder == checkChar
}

func getCharValue(char string) int {
	var result int

	values := make(map[string]int)
	values["A"] = 1
	values["B"] = 2
	values["C"] = 3
	values["D"] = 4
	values["E"] = 5
	values["F"] = 6
	values["G"] = 7
	values["H"] = 8
	values["J"] = 1
	values["K"] = 2
	values["L"] = 3
	values["M"] = 4
	values["N"] = 5
	values["O"] = 6
	values["P"] = 7
	values["R"] = 9
	values["S"] = 2
	values["T"] = 3
	values["U"] = 4
	values["V"] = 5
	values["W"] = 6
	values["X"] = 7
	values["Y"] = 8
	values["Z"] = 9

	if num, ok := values[char]; ok {
		result = num
	} else {
		result, _ = strconv.Atoi(char)
	}

	return result
}

func getVINWeight(position int) int {
	weights := []int{
		8, 7, 6, 5, 4, 3, 2,
		10, 0, // check digit
		9, 8, 7, 6, 5, 4, 3, 2,
	}

	return weights[position]
}
