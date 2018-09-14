package book

import (
	"github.com/louisevanderlith/husk"
)

type servicesTable struct {
	tbl husk.Tabler
}

func NewServicesTable() servicesTable {
	result := husk.NewTable(new(Service))

	return servicesTable{result}
}

func (t servicesTable) Create(obj Service) (serviceRecord, error) {
	set := t.tbl.Create(obj)

	return serviceRecord{set.Record}, set.Error
}

type serviceRecord struct {
	rec husk.Recorder
}
