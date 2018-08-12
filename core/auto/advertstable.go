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

func (t advertsTable) FindByID(id int64) (advertRecord, error) {
	result, err := t.tbl.FindByID(id)

	return advertRecord{result}, err
}

func (t advertsTable) Find(page, pageSize int, filter advertFilter) (advertSet, error) {
	huskFilter, err := husk.MakeFilter(filter)

	var result advertSet

	if err == nil {
		items := t.tbl.Find(page, pageSize, huskFilter)
		result = advertSet{items}
	}

	return result, err
}

func (t advertsTable) FindFirst(filter advertFilter) (advertRecord, error) {
	huskFilter, err := husk.MakeFilter(filter)

	var result husk.Recorder

	if err == nil {
		result, err = t.tbl.FindFirst(huskFilter)
	}

	return advertRecord{result}, err
}

func (t advertsTable) Exists(filter advertFilter) (bool, error) {
	huskFilter, err := husk.MakeFilter(filter)

	result := true

	if err == nil {
		result, err = t.tbl.Exists(huskFilter)
	}

	return result, err
}

func (t advertsTable) Create(obj Advert) (advertRecord, error) {
	result, err := t.tbl.Create(obj)

	return advertRecord{result}, err
}

func (t advertsTable) Update(record advertRecord) error {
	result := t.tbl.Update(record.rec)

	return result
}

func (t advertsTable) Delete(id int64) error {
	return t.tbl.Delete(id)
}

type advertRecord struct {
	rec husk.Recorder
}

func (r advertRecord) Data() *Advert {
	return r.rec.Data().(*Advert)
}

type advertFilter func(o Advert) bool

type advertSet struct {
	*husk.RecordSet
}

func newadvertSet() *advertSet {
	result := husk.NewRecordSet()

	return &advertSet{result}
}
