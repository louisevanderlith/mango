package book

import (
	"github.com/louisevanderlith/husk"
)

type vehiclesTable struct {
	tbl husk.Tabler
}

func NewVehiclesTable() vehiclesTable {
	result := husk.NewTable(new(Vehicle))

	return vehiclesTable{result}
}

func (t vehiclesTable) Create(obj Vehicle) (vehicleRecord, error) {
	result, err := t.tbl.Create(obj)

	return vehicleRecord{result}, err
}

type vehicleRecord struct {
	rec husk.Recorder
}
