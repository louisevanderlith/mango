package things

import "github.com/louisevanderlith/db"

type Manufacturers []*Manufacturer

func (o Manufacturers) Each(handler func(obj db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Manufacturers) Length() int {
	return len(o)
}

func (o Manufacturers) At(index int) db.IRecord {
	return o[index]
}

func (o Manufacturers) Add(obj db.IRecord) {
	item, ok := obj.(*Manufacturer)

	if ok {
		o = append(o, item)
	}
}
