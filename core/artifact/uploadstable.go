package artifact

import (
	"log"

	"github.com/louisevanderlith/husk"
)

type uploadsTable struct {
	tbl husk.Tabler
}

func NewUploadsTable() uploadsTable {
	result := husk.NewTable(new(Upload))

	return uploadsTable{result}
}

func (t uploadsTable) FindByKey(key husk.Key) (uploadRecord, error) {
	result, err := t.tbl.FindByKey(key)

	return uploadRecord{result}, err
}

func (t uploadsTable) Find(page, pageSize int, filter uploadFilter) (uploadSet, error) {
	huskFilter, err := husk.MakeFilter(filter)

	var result uploadSet

	if err == nil {
		items := t.tbl.Find(page, pageSize, huskFilter)
		result = uploadSet{items}
	}

	return result, err
}

func (t uploadsTable) FindFirst(filter uploadFilter) (uploadRecord, error) {
	huskFilter, err := husk.MakeFilter(filter)

	if err != nil {
		return uploadRecord{}, err
	}

	result := t.tbl.FindFirst(huskFilter)

	return uploadRecord{result}, nil
}

func (t uploadsTable) Exists(filter uploadFilter) bool {
	huskFilter, err := husk.MakeFilter(filter)

	if err != nil {
		log.Println(err.Error())
		return true
	}

	return t.tbl.Exists(huskFilter)
}

func (t uploadsTable) Create(obj Upload) (uploadRecord, error) {
	set := t.tbl.Create(obj)
	defer t.tbl.Save()

	return uploadRecord{set.Record}, set.Error
}

func (t uploadsTable) Update(record uploadRecord) error {
	result := t.tbl.Update(record.rec)
	defer t.tbl.Save()

	return result
}

func (t uploadsTable) Delete(key husk.Key) error {
	return t.tbl.Delete(key)
}

func (t uploadsTable) Save() {
	t.tbl.Save()
}

type uploadRecord struct {
	rec husk.Recorder
}

func (r uploadRecord) GetID() int64 {
	return r.GetID()
}

func (r uploadRecord) Data() *Upload {
	return r.rec.Data().(*Upload)
}

type uploadFilter func(o Upload) bool

type uploadSet struct {
	*husk.RecordSet
}

func newUploadSet() *uploadSet {
	result := husk.NewRecordSet()

	return &uploadSet{result}
}
