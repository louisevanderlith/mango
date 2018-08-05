package things

import (
	"github.com/louisevanderlith/husk"
)

type manufacturersTable struct {
	tbl husk.Tabler
}

func NewManufacturersTable() manufacturersTable {
	result := husk.NewTable(new(Manufacturer))

	return manufacturersTable{result}
}

func (t manufacturersTable) Create(obj Manufacturer) (manufacturerRecord, error) {
	result, err := t.tbl.Create(obj)

	return manufacturerRecord{result}, err
}

type manufacturerRecord struct {
	rec husk.Recorder
}
