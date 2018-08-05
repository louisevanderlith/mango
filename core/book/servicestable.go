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
	result, err := t.tbl.Create(obj)

	return serviceRecord{result}, err
}

type serviceRecord struct {
	rec husk.Recorder
}
