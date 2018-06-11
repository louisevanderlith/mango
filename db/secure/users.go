package secure

import "github.com/louisevanderlith/db"

type Users []*User

func (o Users) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Users) Length() int {
	return len(o)
}

func (o Users) At(index int) db.IRecord {
	return o[index]
}

func (o Users) Add(obj db.IRecord) {
	item, ok := obj.(*User)

	if ok {
		o = append(o, item)
	}
}
