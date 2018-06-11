package folio

import "github.com/louisevanderlith/db"

type Abouts []*About

func (o Abouts) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Abouts) Length() int {
	return len(o)
}

func (o Abouts) At(index int) db.IRecord {
	return o[index]
}

func (o Abouts) Add(obj db.IRecord) {
	item, ok := obj.(*About)

	if ok {
		o = append(o, item)
	}
}
