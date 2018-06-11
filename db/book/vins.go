package book

import "github.com/louisevanderlith/db"

type VINs []*VIN

func (o VINs) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o VINs) Length() int {
	return len(o)
}

func (o VINs) At(index int) db.IRecord {
	return o[index]
}

func (o VINs) Add(obj db.IRecord) {
	item, ok := obj.(*VIN)

	if ok {
		o = append(o, item)
	}
}
