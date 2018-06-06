package things

import "github.com/louisevanderlith/db"

type Subcategories []*Subcategory

func (o Subcategories) Each(handler func(obj db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Subcategories) Length() int {
	return len(o)
}

func (o Subcategories) At(index int) db.IRecord {
	return o[index]
}

func (o Subcategories) Add(obj db.IRecord) {
	item, ok := obj.(*Subcategory)

	if ok {
		o = append(o, item)
	}
}
