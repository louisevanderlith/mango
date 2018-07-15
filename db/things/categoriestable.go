package things

import (
	"github.com/louisevanderlith/husk"
)

type categoriesTable struct {
	tbl husk.Tabler
}

func NewCategoriesTable() categoriesTable {
	result := husk.NewTable(new(Category))

	return categoriesTable{result}
}

func (t categoriesTable) Create(obj Category) (categoryRecord, error) {
	result, err := t.tbl.Create(obj)

	return categoryRecord{result}, err
}

type categoryRecord struct {
	rec husk.Recorder
}
