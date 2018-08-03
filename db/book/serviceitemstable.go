package book

import (
	"log"

	"github.com/louisevanderlith/husk"
)

type serviceItemsTable struct {
	tbl husk.Tabler
}

func NewServiceItemsTable() serviceItemsTable {
	result := husk.NewTable(new(ServiceItem))

	return serviceItemsTable{result}
}

func (t serviceItemsTable) FindByID(id int64) (serviceItemRecord, error) {
	result, err := t.tbl.FindByID(id)

	return serviceItemRecord{result}, err
}

func (t serviceItemsTable) Find(page, pageSize int, filter serviceItemFilter) []serviceItemRecord {
	var result []serviceItemRecord
	hskFilter, err := husk.MakeFilter(filter)

	if err != nil {
		log.Print(err)
		return result
	}

	return t.tbl.Find(page, pageSize, hskFilter)
}

func (t serviceItemsTable) FindFirst(filter serviceItemFilter) serviceItemRecord {
	hskFilter, err := husk.MakeFilter(filter)

	if err != nil {
		log.Print(err)
	}

	return serviceItemRecord{t.tbl.FindFirst(hskFilter)}
}

func (t serviceItemsTable) Exists(filter serviceItemFilter) bool {
	result := t.tbl.Exists(filter)

	return result
}

func (t serviceItemsTable) Create(obj ServiceItem) (serviceItemRecord, error) {
	result, err := t.tbl.Create(obj)

	return serviceItemRecord{result}, err
}

func (t serviceItemsTable) Update(record serviceItemRecord) error {
	result := t.tbl.Update(record)

	return result
}

func (t serviceItemsTable) Delete(id int64) error {
	return t.tbl.Delete(id)
}

type serviceItemRecord struct {
	rec husk.Recorder
}

func (r serviceItemRecord) Data() *ServiceItem {
	return r.rec.Data().(*ServiceItem)
}

type serviceItemFilter func(o ServiceItem) bool
