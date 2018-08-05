package secure

import (
	"github.com/louisevanderlith/husk"
)

type loginTracesTable struct {
	tbl husk.Tabler
}

func NewLoginTracesTable() loginTracesTable {
	result := husk.NewTable(new(LoginTrace))

	return loginTracesTable{result}
}

func (t loginTracesTable) Create(obj LoginTrace) (loginTraceRecord, error) {
	result, err := t.tbl.Create(obj)

	return loginTraceRecord{result}, err
}

type loginTraceRecord struct {
	rec husk.Recorder
}
