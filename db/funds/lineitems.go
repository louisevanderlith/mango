package funds

import "github.com/louisevanderlith/db"

type LineItems []*LineItem

func (o LineItems) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o LineItems) Length() int {
	return len(o)
}

func (o LineItems) At(index int) db.IRecord {
	return o[index]
}

func (o LineItems) Add(obj db.IRecord) {
	item, ok := obj.(*LineItem)

	if ok {
		o = append(o, item)
	}
}
