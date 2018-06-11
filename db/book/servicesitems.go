package book

import "github.com/louisevanderlith/db"

type ServiceItems []*ServiceItem

func (o ServiceItems) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o ServiceItems) Length() int {
	return len(o)
}

func (o ServiceItems) At(index int) db.IRecord {
	return o[index]
}

func (o ServiceItems) Add(obj db.IRecord) {
	item, ok := obj.(*ServiceItem)

	if ok {
		o = append(o, item)
	}
}
