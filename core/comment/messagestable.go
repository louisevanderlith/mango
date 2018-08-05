package comment

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

func (t messagesTable) FindByID(id int64) (messageRecord, error) {
	result, err := t.tbl.FindByID(id)

	return messageRecord{result}, err
}

func (t messagesTable) Find(page, pageSize int, filter messageFilter) (messageSet, error) {
	huskFilter, err := husk.MakeFilter(filter)

	var result messageSet

	if err == nil {
		items := t.tbl.Find(page, pageSize, huskFilter)
		result = messageSet{items}
	}

	return result, err
}

func (t messagesTable) FindFirst(filter messageFilter) (messageRecord, error) {
	huskFilter, err := husk.MakeFilter(filter)

	var result husk.Recorder

	if err == nil {
		result, err = t.tbl.FindFirst(huskFilter)
	}

	return messageRecord{result}, err
}

func (t messagesTable) Exists(filter messageFilter) (bool, error) {
	huskFilter, err := husk.MakeFilter(filter)

	result := true

	if err == nil {
		result, err = t.tbl.Exists(huskFilter)
	}

	return result, err
}

func (t messagesTable) Create(obj Message) (messageRecord, error) {
	result, err := t.tbl.Create(obj)

	return messageRecord{result}, err
}

func (t messagesTable) Update(record messageRecord) error {
	result := t.tbl.Update(record.rec)

	return result
}

func (t messagesTable) Delete(id int64) error {
	return t.tbl.Delete(id)
}

type messageRecord struct {
	rec husk.Recorder
}

func (r messageRecord) Data() *Message {
	return r.rec.Data().(*Message)
}

type messageFilter func(o Message) bool

type messageSet struct {
	*husk.RecordSet
}

func newMessageSet() *messageSet {
	result := husk.NewRecordSet()

	return &messageSet{result}
}
