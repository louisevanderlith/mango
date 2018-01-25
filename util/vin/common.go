package vin

import "strconv"

func getWeight(char string, illegalCharacters ...string) int {
	resultWeight := getBaseWeight(char)
	baseWeight := resultWeight
	illegalList := getIllegalMap(illegalCharacters...)

	for _, v := range illegalList {
		if baseWeight > v {
			resultWeight--
		}
	}

	return resultWeight
}

func getBaseWeight(char string) int {
	var result = 0

	if num, err := strconv.Atoi(char); err == nil {
		if num == 0 {
			result = 36
		} else {
			base := 26
			result = base + num
		}
	} else {
		crune := []rune(char)[0]
		result = int(crune) % 32
	}

	return result
}

func getIllegalMap(illegalCharacters ...string) []int {
	var result []int
	keys := make(map[string]int)

	keys["I"] = 9
	keys["O"] = 15
	keys["Q"] = 17
	keys["U"] = 21
	keys["Z"] = 26
	keys["0"] = 36

	for _, v := range illegalCharacters {
		if item, ok := keys[v]; ok {
			result = append(result, item)
		}
	}

	return result
}
