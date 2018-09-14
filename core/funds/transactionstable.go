package funds

import (
	"github.com/louisevanderlith/husk"
)

type transactionsTable struct {
	tbl husk.Tabler
}

func NewTransactionsTable() transactionsTable {
	result := husk.NewTable(new(Requisition))

	return transactionsTable{result}
}

func (t transactionsTable) Create(obj Requisition) (transactionRecord, error) {
	set := t.tbl.Create(obj)

	return transactionRecord{set.Record}, set.Error
}

type transactionRecord struct {
	rec husk.Recorder
}
