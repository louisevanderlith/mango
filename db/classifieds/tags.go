package classifieds

import "github.com/louisevanderlith/db"

type Tags []*Tag

func (o Tags) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Tags) Length() int {
	return len(o)
}

func (o Tags) At(index int) db.IRecord {
	return o[index]
}

func (o Tags) Add(obj db.IRecord) {
	item, ok := obj.(*Tag)

	if ok {
		o = append(o, item)
	}
}
