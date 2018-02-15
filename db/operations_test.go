package db

import (
	"testing"
)

func TestGetFilterValues_Pointers(t *testing.T) {
	input := testTableB{}
	actual := getFilterValues(&input)

	if val, ok := actual["Relation"]; ok {
		t.Error("Pointer has value %s", val)
	}
}

func TestGetRelationships(t *testing.T) {
	input := testTableB{}
	input.Collections = []*testTable{
		&testTable{
			Name: "TEST",
			Age:  8,
		},
	}
	input.Relation = &testTable{
		Name: "RELATE",
		Age:  99,
	}

	actual := getRelationships(&input)

	for _, v := range actual {
		fields := getReadColumns(v)

		if len(fields) <= 0 {
			t.Errorf("no fields found")
		}
	}
}
