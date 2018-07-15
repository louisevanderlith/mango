package funds

import (
	"github.com/louisevanderlith/husk"
)

type heroesTable struct {
	tbl husk.Tabler
}

func NewHeroesTable() heroesTable {
	result := husk.NewTable(new(Hero))

	return heroesTable{result}
}

func (t heroesTable) Create(obj Hero) (heroRecord, error) {
	result, err := t.tbl.Create(obj)

	return heroRecord{result}, err
}

type heroRecord struct {
	rec husk.Recorder
}
