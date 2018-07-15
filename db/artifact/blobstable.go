package artifact

import "github.com/louisevanderlith/husk"

type blobsTable struct {
	tbl husk.Tabler
}

func NewBLOBSTable() blobsTable {
	result := husk.NewTable(new(Blob))

	return blobsTable{result}
}

func (t blobsTable) FindByID(id int64) (blobRecord, error) {
	result, err := t.tbl.FindByID(id)

	return blobRecord{result}, err
}

func (t blobsTable) Find(page, pageSize int, filter blobFilter) []blobRecord {
	result := t.tbl.Find(page, pageSize, filter.ToFilter())

	return blobRecords(result)
}

func (t blobsTable) FindFirst(filter blobFilter) blobRecord {
	result := t.tbl.FindFirst(filter.ToFilter())

	return blobRecord{result}
}

func (t blobsTable) Exists(filter blobFilter) bool {
	result := t.tbl.Exists(filter.ToFilter())

	return result
}

func (t blobsTable) Create(obj Upload) (uploadRecord, error) {
	result, err := t.tbl.Create(obj)

	return uploadRecord{result}, err
}

func (t blobsTable) Update(record uploadRecord) error {
	result := t.tbl.Update(record.rec)

	return result
}

func (t blobsTable) Delete(id int64) error {
	return t.tbl.Delete(id)
}

type blobRecord struct {
	rec husk.Recorder
}

func (r blobRecord) Data() *Blob {
	return r.rec.Data().(*Blob)
}

func blobRecords(records []husk.Recorder) []blobRecord {
	var result []blobRecord

	for _, v := range records {
		result = append(result, blobRecord{v})
	}

	return result
}

type blobFilter func(Blob) bool

func (f blobFilter) ToFilter() husk.Filter {
	return func(obj husk.Dataer) bool {
		return f(obj.(Blob))
	}
}
