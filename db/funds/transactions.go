package funds

import "github.com/louisevanderlith/db"

type Transactions []*Transaction

func (o Transactions) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Transactions) Length() int {
	return len(o)
}

func (o Transactions) At(index int) db.IRecord {
	return o[index]
}

func (o Transactions) Add(obj db.IRecord) {
	item, ok := obj.(*Transaction)

	if ok {
		o = append(o, item)
	}
}
