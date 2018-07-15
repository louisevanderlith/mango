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

func (t usersTable) Create(obj LoginTrace) (userRecord, error) {
	result, err := t.tbl.Create(obj)

	return userRecord{result}, err
}

type userRecord struct {
	rec husk.Recorder
}
