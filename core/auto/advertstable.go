package auto

import (
	"github.com/louisevanderlith/husk"
)

type advertsTable struct {
	tbl husk.Tabler
}

func NewAdvertsTable() advertsTable {
	result := husk.NewTable(new(Advert))

	return advertsTable{result}
}

func (t advertsTable) Create(obj Advert) (advertRecord, error) {
	result, err := t.tbl.Create(obj)

	return advertRecord{result}, err
}

type advertRecord struct {
	rec husk.Recorder
}
