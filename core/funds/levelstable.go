package funds

import (
	"github.com/louisevanderlith/husk"
)

type levelsTable struct {
	tbl husk.Tabler
}

func NewLevelsTable() levelsTable {
	result := husk.NewTable(new(Level))

	return levelsTable{result}
}

func (t levelsTable) Create(obj Level) (levelRecord, error) {
	result, err := t.tbl.Create(obj)

	return levelRecord{result}, err
}

type levelRecord struct {
	rec husk.Recorder
}
