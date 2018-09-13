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

func (t heroesTable) FindByKey(key husk.Key) (heroRecord, error) {
	result, err := t.tbl.FindByKey(key)

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

func (t heroesTable) Create(obj Hero) (heroRecord, error) {
	set := t.tbl.Create(obj)

	return heroRecord{set.Record}, set.Error
}

func (t heroesTable) Update(record heroRecord) error {
	result := t.tbl.Update(record.rec)

	return result
}

func (t heroesTable) Delete(key husk.Key) error {
	return t.tbl.Delete(key)
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
