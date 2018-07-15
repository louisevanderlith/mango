package folio

import (
	"github.com/louisevanderlith/husk"
)

type profilesTable struct {
	tbl husk.Tabler
}

func NewProfilesTable() profilesTable {
	result := husk.NewTable(new(Profile))

	return profilesTable{result}
}

func (t profilesTable) Create(obj Portfolio) (profileRecord, error) {
	result, err := t.tbl.Create(obj)

	return profileRecord{result}, err
}

type profileRecord struct {
	rec husk.Recorder
}
