package funds

import (
	"github.com/louisevanderlith/husk"
)

type experiencesTable struct {
	tbl husk.Tabler
}

func NewExperiencesTable() experiencesTable {
	result := husk.NewTable(new(Experience))

	return experiencesTable{result}
}

func (t experiencesTable) Create(obj Experience) (experienceRecord, error) {
	result, err := t.tbl.Create(obj)

	return experienceRecord{result}, err
}

type experienceRecord struct {
	rec husk.Recorder
}
