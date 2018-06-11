package book

import "github.com/louisevanderlith/db"

type Vehicles []*Vehicle

func (o Vehicles) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Vehicles) Length() int {
	return len(o)
}

func (o Vehicles) At(index int) db.IRecord {
	return o[index]
}

func (o Vehicles) Add(obj db.IRecord) {
	item, ok := obj.(*Vehicle)

	if ok {
		o = append(o, item)
	}
}
