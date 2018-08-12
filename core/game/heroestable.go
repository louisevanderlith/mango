package game

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

func (t heroesTable) FindByID(id int64) (heroRecord, error) {
	result, err := t.tbl.FindByID(id)

	return heroRecord{result}, err
}

func (t heroesTable) Find(page, pageSize int, filter heroFilter) (heroSet, error) {
	huskFilter, err := husk.MakeFilter(filter)

	var result heroSet

	if err == nil {
		items := t.tbl.Find(page, pageSize, huskFilter)
		result = heroSet{items}
	}

	return result, err
}

func (t heroesTable) FindFirst(filter heroFilter) (heroRecord, error) {
	huskFilter, err := husk.MakeFilter(filter)

	var result husk.Recorder

	if err == nil {
		result, err = t.tbl.FindFirst(huskFilter)
	}

	return heroRecord{result}, err
}

func (t heroesTable) Exists(filter heroFilter) (bool, error) {
	huskFilter, err := husk.MakeFilter(filter)

	result := true

	if err == nil {
		result, err = t.tbl.Exists(huskFilter)
	}

	return result, err
}

func (t heroesTable) Create(obj hero) (heroRecord, error) {
	result, err := t.tbl.Create(obj)

	return heroRecord{result}, err
}

func (t heroesTable) Update(record heroRecord) error {
	result := t.tbl.Update(record.rec)

	return result
}

func (t heroesTable) Delete(id int64) error {
	return t.tbl.Delete(id)
}

type heroRecord struct {
	rec husk.Recorder
}

func (r heroRecord) Data() *Hero {
	return r.rec.Data().(*Hero)
}

type heroFilter func(o Hero) bool

type heroSet struct {
	*husk.RecordSet
}

func newheroSet() *heroSet {
	result := husk.NewRecordSet()

	return &heroSet{result}
}
