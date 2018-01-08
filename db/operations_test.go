package db

import (
	"testing"
)

func TestgetFilterValues_Pointers(t *testing.T) {
	input := testTableB{}
	actual := getFilterValues(input)

	if val, ok := actual["Relation"]; ok {
		t.Error("Pointer has value %s", val)
	}
}
