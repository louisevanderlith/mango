package funds

import (
	"github.com/louisevanderlith/husk"
)

type lineItemsTable struct {
	tbl husk.Tabler
}

func NewLineItemsTable() lineItemsTable {
	result := husk.NewTable(new(LineItem))

	return lineItemsTable{result}
}

func (t lineItemsTable) Create(obj Hero) (lineItemRecord, error) {
	result, err := t.tbl.Create(obj)

	return lineItemRecord{result}, err
}

type lineItemRecord struct {
	rec husk.Recorder
}
