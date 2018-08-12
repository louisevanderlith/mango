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

func (t portfoliosTable) FindByID(id int64) (portfolioRecord, error) {
	result, err := t.tbl.FindByID(id)

	return portfolioRecord{result}, err
}

func (t portfoliosTable) Find(page, pageSize int, filter portfolioFilter) (portfolioSet, error) {
	huskFilter, err := husk.MakeFilter(filter)

	var result portfolioSet

	if err == nil {
		items := t.tbl.Find(page, pageSize, huskFilter)
		result = portfolioSet{items}
	}

	return result, err
}

func (t portfoliosTable) FindFirst(filter portfolioFilter) (portfolioRecord, error) {
	huskFilter, err := husk.MakeFilter(filter)

	var result husk.Recorder

	if err == nil {
		result, err = t.tbl.FindFirst(huskFilter)
	}

	return portfolioRecord{result}, err
}

func (t portfoliosTable) Exists(filter portfolioFilter) (bool, error) {
	huskFilter, err := husk.MakeFilter(filter)

	result := true

	if err == nil {
		result, err = t.tbl.Exists(huskFilter)
	}

	return result, err
}

func (t portfoliosTable) Create(obj Portfolio) (portfolioRecord, error) {
	result, err := t.tbl.Create(obj)

	return portfolioRecord{result}, err
}

func (t portfoliosTable) Update(record portfolioRecord) error {
	result := t.tbl.Update(record.rec)

	return result
}

func (t portfoliosTable) Delete(id int64) error {
	return t.tbl.Delete(id)
}

type portfolioRecord struct {
	rec husk.Recorder
}

func (r portfolioRecord) Data() *Portfolio {
	return r.rec.Data().(*Portfolio)
}

func (r *portfolioRecord) Set(portfolio Portfolio) error

type portfolioFilter func(o Portfolio) bool

type portfolioSet struct {
	*husk.RecordSet
}

func newportfolioSet() *portfolioSet {
	result := husk.NewRecordSet()

	return &portfolioSet{result}
}
