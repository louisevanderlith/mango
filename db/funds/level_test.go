package funds

import "testing"

func init() {
	seedLevel()
}

func TestXPForLevel(t *testing.T) {
	input := 68
	expected := 117300
	actual := xpRequired(input)

	if expected != actual {
		t.Errorf("Expected: %v, got %v", expected, actual)
	}
}

func TestLevelForXP(t *testing.T) {
	input := 117299
	expected := 67
	actual := getLevel(input)

	if expected != actual.Rank {
		t.Errorf("Expected: %v, got %v", expected, actual)
	}
}
