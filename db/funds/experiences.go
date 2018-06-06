package funds

import "github.com/louisevanderlith/db"

type Experiences []*Experience

func (o Experiences) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Experiences) Length() int {
	return len(o)
}

func (o Experiences) At(index int) db.IRecord {
	return o[index]
}

func (o Experiences) Add(obj db.IRecord) {
	item, ok := obj.(*Experience)

	if ok {
		o = append(o, item)
	}
}
