package db

import (
	"reflect"
)

type Set struct {
	t     reflect.Type
	items map[int64]IRecord
}

func NewSet(t IRecord) *Set {
	result := &Set{}
	result.t = reflect.TypeOf(t)
	result.items = make(map[int64]IRecord)

	return result
}

func (set *Set) Create(item IRecord) (id int64, err error) {
	if reflect.TypeOf(item) == set.t {

		var valid bool
		valid, err = item.Validate()

		if valid {
			var exists bool
			exists, err = item.Exists()

			if !exists {
				id, err = insert(item)
				set.items[id] = item
			}
		}
	}

	return id, err
}

func (set *Set) ReadOne(filter IRecord) (IRecord, error) {

	err := read(filter)

	if err == nil {
		_, ok := set.items[filter.GetID()]

		if !ok {
			set.items[filter.GetID()] = filter
		}
	}

	return filter, err
}

func (set *Set) Read(filter IRecord, container interface{}) error {
	err := readAll(filter, container)

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

func (set *Set) Update(item IRecord) {
	id := item.GetID()
	_, ok := set.items[id]

	if ok {
		set.items[id] = item
		update(item)
	}
}

func (set *Set) Delete(item IRecord) {
	id := item.GetID()
	memItem, ok := set.items[id]

	if ok {
		update(memItem.Disable())
	}
}
