package comms

import (
	"github.com/louisevanderlith/husk"
)

type messagesTable struct {
	tbl husk.Tabler
}

func NewMessagesTable() messagesTable {
	result := husk.NewTable(new(Message))

	return messagesTable{result}
}

func (t messagesTable) Create(obj Message) (messageRecord, error) {
	result, err := t.tbl.Create(obj)

	return messageRecord{result}, err
}

type messageRecord struct {
	rec husk.Recorder
}
