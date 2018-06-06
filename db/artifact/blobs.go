package artifact

import "github.com/louisevanderlith/db"

type Blobs []*Blob

func (o Blobs) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Blobs) Length() int {
	return len(o)
}

func (o Blobs) At(index int) db.IRecord {
	return o[index]
}

func (o Blobs) Add(obj db.IRecord) {
	item, ok := obj.(*Blob)

	if ok {
		o = append(o, item)
	}
}
