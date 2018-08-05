package folio

import (
	"github.com/louisevanderlith/husk"
)

type portfoliosTable struct {
	tbl husk.Tabler
}

func NewPortfoliosTable() portfoliosTable {
	result := husk.NewTable(new(Portfolio))

	return portfoliosTable{result}
}

func (t portfoliosTable) Create(obj Portfolio) (portfolioRecord, error) {
	result, err := t.tbl.Create(obj)

	return portfolioRecord{result}, err
}

type portfolioRecord struct {
	rec husk.Recorder
}
