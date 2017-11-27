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

		id, err = insert(item)

		set.items[id] = item
	}

	return
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

func (set *Set) Read(filter IRecord, container *[]*IRecord) error {
	err := readAll(filter, container)

	if err == nil {
		for _, v := range *container {
			item := *v
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
		memItem = memItem.Disable()
		update(memItem)
	}
}
