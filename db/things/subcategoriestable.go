package things

import (
	"github.com/louisevanderlith/husk"
)

type subcategoriesTable struct {
	tbl husk.Tabler
}

func NewSubcategoriesTable() subcategoriesTable {
	result := husk.NewTable(new(Subcategory))

	return subcategoriesTable{result}
}

func (t subcategoriesTable) Create(obj Subcategory) (subcategoryRecord, error) {
	result, err := t.tbl.Create(obj)

	return subcategoryRecord{result}, err
}

type subcategoryRecord struct {
	rec husk.Recorder
}
