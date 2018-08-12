package secure

import (
	"github.com/louisevanderlith/husk"
)

type usersTable struct {
	tbl husk.Tabler
}

func NewUsersTable() usersTable {
	result := husk.NewTable(new(User))

	return usersTable{result}
}

func (t usersTable) FindByID(id int64) (userRecord, error) {
	result, err := t.tbl.FindByID(id)

	return userRecord{result}, err
}

func (t usersTable) Find(page, pageSize int, filter userFilter) (userSet, error) {
	huskFilter, err := husk.MakeFilter(filter)

	var result userSet

	if err == nil {
		items := t.tbl.Find(page, pageSize, huskFilter)
		result = userSet{items}
	}

	return result, err
}

func (t usersTable) FindFirst(filter userFilter) (userRecord, error) {
	huskFilter, err := husk.MakeFilter(filter)

	var result husk.Recorder

	if err == nil {
		result, err = t.tbl.FindFirst(huskFilter)
	}

	return userRecord{result}, err
}

func (t usersTable) Exists(filter userFilter) (bool, error) {
	huskFilter, err := husk.MakeFilter(filter)

	result := true

	if err == nil {
		result, err = t.tbl.Exists(huskFilter)
	}

	return result, err
}

func (t usersTable) Create(obj *User) (userRecord, error) {
	result, err := t.tbl.Create(obj)

	return userRecord{result}, err
}

func (t usersTable) Update(record userRecord) error {
	result := t.tbl.Update(record.rec)

	return result
}

func (t usersTable) Delete(id int64) error {
	return t.tbl.Delete(id)
}

type userRecord struct {
	rec husk.Recorder
}

func (r userRecord) Data() *User {
	return r.rec.Data().(*User)
}

type userFilter func(o User) bool

type userSet struct {
	*husk.RecordSet
}

func newUserSet() *userSet {
	result := husk.NewRecordSet()

	return &userSet{result}
}
