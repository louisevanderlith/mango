package secure

import "github.com/louisevanderlith/db"

type LoginTraces []*LoginTrace

func (o LoginTraces) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o LoginTraces) Length() int {
	return len(o)
}

func (o LoginTraces) At(index int) db.IRecord {
	return o[index]
}

func (o LoginTraces) Add(obj db.IRecord) {
	item, ok := obj.(*LoginTrace)

	if ok {
		o = append(o, item)
	}
}
