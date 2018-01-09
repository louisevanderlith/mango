package db

import (
	"testing"
)

func TestGetFilterValues_Pointers(t *testing.T) {
	input := testTableB{}
	actual := getFilterValues(input)

	if val, ok := actual["Relation"]; ok {
		t.Error("Pointer has value %s", val)
	}
}