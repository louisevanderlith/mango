package folio

import "github.com/louisevanderlith/db"

type Headers []*Header

func (o Headers) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Headers) Length() int {
	return len(o)
}

func (o Headers) At(index int) db.IRecord {
	return o[index]
}

func (o Headers) Add(obj db.IRecord) {
	item, ok := obj.(*Header)

	if ok {
		o = append(o, item)
	}
}
