package book

import "github.com/louisevanderlith/db"

type Services []*Service

func (o Services) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Services) Length() int {
	return len(o)
}

func (o Services) At(index int) db.IRecord {
	return o[index]
}

func (o Services) Add(obj db.IRecord) {
	item, ok := obj.(*Service)

	if ok {
		o = append(o, item)
	}
}
