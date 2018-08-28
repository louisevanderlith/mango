package artifact

import "github.com/louisevanderlith/husk"

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

	var result husk.Recorder

	if err == nil {
		result, err = t.tbl.FindFirst(huskFilter)
	}

	return uploadRecord{result}, err
}

func (t uploadsTable) Exists(filter uploadFilter) (bool, error) {
	huskFilter, err := husk.MakeFilter(filter)

	result := true

	if err == nil {
		result, err = t.tbl.Exists(huskFilter)
	}

	return result, err
}

func (t uploadsTable) Create(obj Upload) (uploadRecord, error) {
	result, err := t.tbl.Create(obj)

	return uploadRecord{result}, err
}

func (t uploadsTable) Update(record uploadRecord) error {
	result := t.tbl.Update(record.rec)

	return result
}

func (t uploadsTable) Delete(key husk.Key) error {
	return t.tbl.Delete(key)
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
