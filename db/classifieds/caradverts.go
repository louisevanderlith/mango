package classifieds

import "github.com/louisevanderlith/db"

type CarAdverts []*CarAdvert

func (o CarAdverts) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o CarAdverts) Length() int {
	return len(o)
}

func (o CarAdverts) At(index int) db.IRecord {
	return o[index]
}

func (o CarAdverts) Add(obj db.IRecord) {
	item, ok := obj.(*CarAdvert)

	if ok {
		o = append(o, item)
	}
}
