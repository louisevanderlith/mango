package things

import "github.com/louisevanderlith/db"

type Categories []*Category

func (o Categories) Each(handler func(obj db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Categories) Length() int {
	return len(o)
}

func (o Categories) At(index int) db.IRecord {
	return o[index]
}

func (o Categories) Add(obj db.IRecord) {
	item, ok := obj.(*Category)

	if ok {
		o = append(o, item)
	}
}
