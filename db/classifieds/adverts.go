package classifieds

import "github.com/louisevanderlith/db"

type Adverts []*Advert

func (o Adverts) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Adverts) Length() int {
	return len(o)
}

func (o Adverts) At(index int) db.IRecord {
	return o[index]
}

func (o Adverts) Add(obj db.IRecord) {
	item, ok := obj.(*Advert)

	if ok {
		o = append(o, item)
	}
}
