package util

import (
	"testing"
)

func TestIsValidPrice_MustBeValid(t *testing.T) {
	inputPrice := 21
	recommendedPrice := 25

	isValid, msg := PriceInBounds(inputPrice, recommendedPrice)

	if !isValid {
		t.Error(msg)
	}
}

func TestIsValidPrice_MustBeInvalid(t *testing.T) {
	inputPrice := 85
	recommendedPrice := 25

	isValid, msg := PriceInBounds(inputPrice, recommendedPrice)

	if !isValid {
		t.Error(msg)
	}
}

func TestGetFraction_ShouldBe14(t *testing.T) {
	inputPrice := 1296
	expectedFraction := 14

	actualFraction := getFraction(inputPrice)

	if expectedFraction != actualFraction {
		t.Errorf("Fraction was not valid, got: %d, want: %d", actualFraction, expectedFraction)
	}
}

func TestGetFraction_ShouldBe9(t *testing.T) {
	inputPrice := 171
	expectedFraction := 9

	actualFraction := getFraction(inputPrice)

	if expectedFraction != actualFraction {
		t.Errorf("Fraction was not valid, got: %d, want: %d", actualFraction, expectedFraction)
	}
}

func TestGetFraction_ShouldBe2(t *testing.T) {
	inputPrice := 7
	expectedFraction := 2

	actualFraction := getFraction(inputPrice)

	if expectedFraction != actualFraction {
		t.Errorf("Fraction was not valid, got: %d, want: %d", actualFraction, expectedFraction)
	}
}

func TestGetFraction_ShouldBe9Between(t *testing.T) {
	inputPrice := 245
	expectedFraction := 9

	actualFraction := getFraction(inputPrice)

	if expectedFraction != actualFraction {
		t.Errorf("Fraction was not valid, got: %d, want: %d", actualFraction, expectedFraction)
	}
}
