package artifact

import "github.com/louisevanderlith/db"

type Uploads []*Upload

func (o Uploads) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Uploads) Length() int {
	return len(o)
}

func (o Uploads) At(index int) db.IRecord {
	return o[index]
}

func (o Uploads) Add(obj db.IRecord) {
	item, ok := obj.(*Upload)

	if ok {
		o = append(o, item)
	}
}
