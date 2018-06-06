package things

import (
	"github.com/louisevanderlith/db"
)

type Models []*Model

func (o Models) Each(handler func(obj db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Models) Length() int {
	return len(o)
}

func (o Models) At(index int) db.IRecord {
	return o[index]
}

func (o Models) Add(obj db.IRecord) {
	item, ok := obj.(*Model)

	if ok {
		o = append(o, item)
	}
}
