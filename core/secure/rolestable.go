package secure

import (
	"github.com/louisevanderlith/husk"
)

type rolesTable struct {
	tbl husk.Tabler
}

func NewRolesTable() rolesTable {
	result := husk.NewTable(new(Role))

	return rolesTable{result}
}

func (t rolesTable) Create(obj Role) (roleRecord, error) {
	result, err := t.tbl.Create(obj)

	return roleRecord{result}, err
}

type roleRecord struct {
	rec husk.Recorder
}
