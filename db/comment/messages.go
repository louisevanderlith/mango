package comment

import "github.com/louisevanderlith/db"

type Messages []*Message

func (o Messages) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Messages) Length() int {
	return len(o)
}

func (o Messages) At(index int) db.IRecord {
	return o[index]
}

func (o Messages) Add(obj db.IRecord) {
	item, ok := obj.(*Message)

	if ok {
		o = append(o, item)
	}
}
