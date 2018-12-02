package book

import (
	"github.com/louisevanderlith/husk"
)

type vinsTable struct {
	tbl husk.Tabler
}

func NewVINSTable() vinsTable {
	result := husk.NewTable(new(VIN))

	return vinsTable{result}
}

func (t vinsTable) Create(obj VIN) (vinRecord, error) {
	set := t.tbl.Create(obj)

	return vinRecord{set.Record}, set.Error
}

type vinRecord struct {
	rec husk.Recorder
}
