package comms

import (
	"log"

	"github.com/louisevanderlith/husk"
)

type messagesTable struct {
	tbl husk.Tabler
}

func NewMessagesTable() messagesTable {
	result := husk.NewTable(new(Message))

	return messagesTable{result}
}

func (t messagesTable) FindByKey(key husk.Key) (messageRecord, error) {
	result, err := t.tbl.FindByKey(key)

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

	if err != nil {
		return messageRecord{}, err
	}

	result := t.tbl.FindFirst(huskFilter)

	return messageRecord{result}, err
}

func (t messagesTable) Exists(filter messageFilter) bool {
	huskFilter, err := husk.MakeFilter(filter)

	if err != nil {
		log.Println(err)
		return true
	}

	return t.tbl.Exists(huskFilter)
}

func (t messagesTable) Create(obj Message) (messageRecord, error) {
	set := t.tbl.Create(obj)
	defer t.tbl.Save()

	return messageRecord{set.Record}, set.Error
}

func (t messagesTable) Update(record messageRecord) error {
	result := t.tbl.Update(record.rec)
	defer t.tbl.Save()

	return result
}

func (t messagesTable) Delete(key husk.Key) error {
	return t.tbl.Delete(key)
}

type messageRecord struct {
	rec husk.Recorder
}

func (r messageRecord) GetID() int64 {
	return r.GetID()
}

func (r messageRecord) Data() *Message {
	return r.rec.Data().(*Message)
}

type messageFilter func(o Message) bool

type messageSet struct {
	*husk.RecordSet
}

func newmessageSet() *messageSet {
	result := husk.NewRecordSet()

	return &messageSet{result}
}
