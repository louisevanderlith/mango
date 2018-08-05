package auto

import (
	"github.com/louisevanderlith/husk"
)

type tagsTable struct {
	tbl husk.Tabler
}

func NewTagsTable() tagsTable {
	result := husk.NewTable(new(Tag))

	return tagsTable{result}
}

func (t tagsTable) Create(obj Tag) (tagRecord, error) {
	result, err := t.tbl.Create(obj)

	return tagRecord{result}, err
}

type tagRecord struct {
	rec husk.Recorder
}
