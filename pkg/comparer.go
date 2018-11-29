package util

import (
	"fmt"
)

// PriceInBounds will check if the requested price is within range of the recommended price
func PriceInBounds(requestedPrice int, recommendedPrice int) (bool, string) {
	fraction := getFraction(recommendedPrice)
	variance := getVariance(recommendedPrice, fraction)
	upperLimit := getUpperLimit(recommendedPrice, variance)
	lowerLimit := getLowerLimit(recommendedPrice, variance)

	result := true
	message := ""

	if requestedPrice > upperLimit || requestedPrice < lowerLimit {
		result = false
		message = fmt.Sprintf("Requested Price isn't between %d and %d", lowerLimit, upperLimit)
	}

	return result, message
}

func getFraction(price int) int {
	fraction := 2
	variance := price

	for variance > 10 {
		variance = (variance / 3) * 2
		fraction++
	}

	return fraction
}

func getVariance(price int, fraction int) int {
	return price / fraction
}

func getUpperLimit(price int, variance int) int {
	return price + variance
}

func getLowerLimit(price int, variance int) int {
	return price - variance
}
