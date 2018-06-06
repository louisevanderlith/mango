package folio

import "github.com/louisevanderlith/db"

type Profiles []*Profile

func (o Profiles) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Profiles) Length() int {
	return len(o)
}

func (o Profiles) At(index int) db.IRecord {
	return o[index]
}

func (o Profiles) Add(obj db.IRecord) {
	item, ok := obj.(*Profile)

	if ok {
		o = append(o, item)
	}
}
