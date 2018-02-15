package db

import (
	"errors"
	"fmt"
	"reflect"
)

type MemSet struct {
	t     reflect.Type
	items map[int64]IRecord
}

// NewSet creates an instance of a Set, which enables type checking and keeping track of records
// # .NewSet(obj.Record{})
func NewMemSet(t IRecord) *MemSet {
	result := &MemSet{}
	result.t = reflect.TypeOf(t)
	result.items = make(map[int64]IRecord)

	return result
}

// Create validates the model and then saves that record to the database.
// # id, err := folio.Ctx.Profile.Create(&profile)
func (set *MemSet) Create(item IRecord) (id int64, err error) {
	t := reflect.TypeOf(item)
	elem := t.Elem()

	if elem == set.t {
		var valid bool
		valid, err = item.Validate()

		if valid {
			var exists bool
			exists, err = item.Exists()

			if !exists {
				set.items[id] = item
			}
		}
	} else {
		msg := fmt.Sprintf("%s is not of type %s", elem, set.t)
		err = errors.New(msg)
	}

	return id, err
}

// CreateMulti inserts multiple records, without running validation
// # manufacturers := []Manufacturer{}
// # count, err := folio.Ctx.Manufacturer.CreateMulti(manufacturers)
func (set *MemSet) CreateMulti(count int, items interface{}) (insertCount int64, err error) {
	return insertMulti(count, items)
}

// ReadOne reads a single record from the database
// filter: An object that has the fields populated that you want to filter on (Filters will always be 'AND')
// related: Relationships are lazy-loaded, to include nested items you must specify them.
// # record, err := testCtx.Profile.ReadOne(Profile{ID: 56}, "User")
func (set *MemSet) ReadOne(filter IRecord, related ...string) (IRecord, error) {

	err := read(filter, related...)

	if err == nil {
		_, ok := set.items[filter.GetID()]

		if !ok {
			set.items[filter.GetID()] = filter
		}
	}

	return filter, err
}

// Read reads all records that fit the filter provided
// filter: An object that has the fields populated that you want to filter on (Filters will always be 'AND')
// container: The result set will populate the container.
// # var results []*artifact.Upload
// # upl := artifact.Upload{Type: "JPEG"}
// # err := artifact.Ctx.Upload.Read(upl, &results)
func (set *MemSet) Read(filter IRecord, container interface{}) error {
	err := readAll(&filter, container)

	if err == nil {
		records := reflect.ValueOf(container).Elem()
		for i := 0; i < records.Len(); i++ {
			item := records.Index(i).Interface().(IRecord)

			_, ok := set.items[item.GetID()]

			if !ok {
				set.items[item.GetID()] = item
			}
		}
	}

	return err
}

// Update saves the provided record to the database.
// The record must exist in the database.
// item: The record you want to update.
// #  testCtx.TableA.Update(row)
func (set *MemSet) Update(item IRecord) error {
	var err error
	id := item.GetID()
	_, ok := set.items[id]

	if ok {
		set.items[id] = item
	} else {
		err = errors.New("item not found in set")
	}

	return err
}

// Delete will delete the record from the database.
// This function currently only deletes a record based on the provided ID
// item: The record containing the ID you want to delete.
// # row := TableA{ID: 99}
// # testCtx.TableA.Delete(row)
func (set *MemSet) Delete(item IRecord) {
	id := item.GetID()
	memItem, _ := set.items[id]
	memItem.Disable()
}
