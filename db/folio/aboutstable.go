package folio

import (
	"github.com/louisevanderlith/husk"
)

type aboutsTable struct {
	tbl husk.Tabler
}

func NewAboutsTable() aboutsTable {
	result := husk.NewTable(new(About))

	return aboutsTable{result}
}

func (t aboutsTable) Create(obj About) (aboutRecord, error) {
	result, err := t.tbl.Create(obj)

	return aboutRecord{result}, err
}

type aboutRecord struct {
	rec husk.Recorder
}
