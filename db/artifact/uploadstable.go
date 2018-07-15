package artifact

import "github.com/louisevanderlith/husk"

type uploadsTable struct {
	tbl husk.Tabler
}

func NewUploadsTable() uploadsTable {
	result := husk.NewTable(new(Upload))

	return uploadsTable{result}
}

func (t uploadsTable) FindByID(id int64) (uploadRecord, error) {
	result, err := t.tbl.FindByID(id)

	return uploadRecord{result}, err
}

func (t uploadsTable) Find(page, pageSize int, filter uploadFilter) []uploadRecord {
	result := t.tbl.Find(page, pageSize, filter)

	return result
}

func (t uploadsTable) FindFirst(filter uploadFilter) uploadRecord {
	result := t.tbl.FindFirst(filter)

	return result
}

func (t uploadsTable) Exists(filter uploadFilter) bool {
	result := t.tbl.Exists(filter)

	return result
}

func (t uploadsTable) Create(obj Upload) (uploadRecord, error) {
	result, err := t.tbl.Create(obj)

	return uploadRecord{result}, err
}

func (t uploadsTable) Update(record uploadRecord) error {
	result := t.tbl.Update(record)

	return result
}

func (t uploadsTable) Delete(id int64) error {
	return t.tbl.Delete(id)
}

type uploadRecord struct {
	rec husk.Recorder
}

func (r uploadRecord) Data() *Upload {
	return r.rec.Data().(*Upload)
}

type uploadFilter func(o Upload) bool
