package funds

import "github.com/louisevanderlith/db"

type Heroes []*Hero

func (o Heroes) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Heroes) Length() int {
	return len(o)
}

func (o Heroes) At(index int) db.IRecord {
	return o[index]
}

func (o Heroes) Add(obj db.IRecord) {
	item, ok := obj.(*Hero)

	if ok {
		o = append(o, item)
	}
}
