package classifieds

import (
	"github.com/louisevanderlith/husk"
)

type carAdvertsTable struct {
	tbl husk.Tabler
}

func NewCarAdvertsTable() carAdvertsTable {
	result := husk.NewTable(new(CarAdvert))

	return carAdvertsTable{result}
}

func (t carAdvertsTable) Create(obj CarAdvert) (carAdvertRecord, error) {
	result, err := t.tbl.Create(obj)

	return carAdvertRecord{result}, err
}

type carAdvertRecord struct {
	rec husk.Recorder
}
