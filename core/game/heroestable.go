package game

import (
	"log"

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

	if err != nil {
		return heroRecord{}, err
	}

	result := t.tbl.FindFirst(huskFilter)

	return heroRecord{result}, err
}

func (t heroesTable) Exists(filter heroFilter) bool {
	huskFilter, err := husk.MakeFilter(filter)

	if err != nil {
		log.Println(err)
		return true
	}

	return t.tbl.Exists(huskFilter)
}

func (t heroesTable) Create(obj Hero) (heroRecord, error) {
	set := t.tbl.Create(obj)
	defer t.tbl.Save()

	return heroRecord{set.Record}, set.Error
}

func (t heroesTable) Update(record heroRecord) error {
	result := t.tbl.Update(record.rec)
	defer t.tbl.Save()

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
