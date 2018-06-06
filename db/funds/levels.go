package funds

import "github.com/louisevanderlith/db"

type Levels []*Level

func (o Levels) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Levels) Length() int {
	return len(o)
}

func (o Levels) At(index int) db.IRecord {
	return o[index]
}

func (o Levels) Add(obj db.IRecord) {
	item, ok := obj.(*Level)

	if ok {
		o = append(o, item)
	}
}
