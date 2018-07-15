package funds

import (
	"github.com/louisevanderlith/husk"
)

type requisitionsTable struct {
	tbl husk.Tabler
}

func NewRequisitionsTable() requisitionsTable {
	result := husk.NewTable(new(Requisition))

	return requisitionsTable{result}
}

func (t requisitionsTable) Create(obj Requisition) (requisitionRecord, error) {
	result, err := t.tbl.Create(obj)

	return requisitionRecord{result}, err
}

type requisitionRecord struct {
	rec husk.Recorder
}
