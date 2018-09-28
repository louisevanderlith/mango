package auto

import (
	"log"

	"github.com/louisevanderlith/husk"
)

type advertsTable struct {
	tbl husk.Tabler
}

func NewAdvertsTable() advertsTable {
	result := husk.NewTable(new(Advert))

	return advertsTable{result}
}

func (t advertsTable) FindByKey(key husk.Key) (advertRecord, error) {
	result, err := t.tbl.FindByKey(key)

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

	if err != nil {
		return advertRecord{}, err
	}

	result := t.tbl.FindFirst(huskFilter)

	return advertRecord{result}, nil
}

func (t advertsTable) Exists(filter advertFilter) bool {
	huskFilter, err := husk.MakeFilter(filter)

	if err != nil {
		log.Println(err)
		return true
	}

	return t.tbl.Exists(huskFilter)
}

func (t advertsTable) Create(obj Advert) (advertRecord, error) {
	set := t.tbl.Create(obj)
	defer t.tbl.Save()

	return advertRecord{set.Record}, set.Error
}

func (t advertsTable) Update(record advertRecord) error {
	result := t.tbl.Update(record.rec)
	defer t.tbl.Save()

	return result
}

func (t advertsTable) Delete(key husk.Key) error {
	return t.tbl.Delete(key)
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
