package folio

import (
	"github.com/louisevanderlith/husk"
)

type headersTable struct {
	tbl husk.Tabler
}

func NewHeadersTable() headersTable {
	result := husk.NewTable(new(Header))

	return headersTable{result}
}

func (t headersTable) Create(obj Header) (headerRecord, error) {
	result, err := t.tbl.Create(obj)

	return headerRecord{result}, err
}

type headerRecord struct {
	rec husk.Recorder
}
