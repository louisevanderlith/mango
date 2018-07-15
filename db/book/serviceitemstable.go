package book

import (
	"github.com/louisevanderlith/husk"
)

type serviceItemsTable struct {
	tbl husk.Tabler
}

func NewServiceItemsTable() serviceItemsTable {
	result := husk.NewTable(new(ServiceItem))

	return serviceItemsTable{result}
}

func (t serviceItemsTable) Create(obj ServiceItem) (serviceItemRecord, error) {
	result, err := t.tbl.Create(obj)

	return serviceItemRecord{result}, err
}

type serviceItemRecord struct {
	rec husk.Recorder
}
