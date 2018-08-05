package things

import (
	"github.com/louisevanderlith/husk"
)

type modelsTable struct {
	tbl husk.Tabler
}

func NewModelsTable() modelsTable {
	result := husk.NewTable(new(Model))

	return modelsTable{result}
}

func (t modelsTable) Create(obj Model) (modelRecord, error) {
	result, err := t.tbl.Create(obj)

	return modelRecord{result}, err
}

type modelRecord struct {
	rec husk.Recorder
}
