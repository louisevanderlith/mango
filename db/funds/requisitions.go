package funds

import "github.com/louisevanderlith/db"

type Requisitions []*Requisition

func (o Requisitions) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Requisitions) Length() int {
	return len(o)
}

func (o Requisitions) At(index int) db.IRecord {
	return o[index]
}

func (o Requisitions) Add(obj db.IRecord) {
	item, ok := obj.(*Requisition)

	if ok {
		o = append(o, item)
	}
}
