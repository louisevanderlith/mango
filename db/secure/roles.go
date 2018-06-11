package secure

import "github.com/louisevanderlith/db"

type Roles []*Role

func (o Roles) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Roles) Length() int {
	return len(o)
}

func (o Roles) At(index int) db.IRecord {
	return o[index]
}

func (o Roles) Add(obj db.IRecord) {
	item, ok := obj.(*Role)

	if ok {
		o = append(o, item)
	}
}
