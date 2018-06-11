package folio

import "github.com/louisevanderlith/db"

type Portfolios []*Portfolio

func (o Portfolios) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o Portfolios) Length() int {
	return len(o)
}

func (o Portfolios) At(index int) db.IRecord {
	return o[index]
}

func (o Portfolios) Add(obj db.IRecord) {
	item, ok := obj.(*Portfolio)

	if ok {
		o = append(o, item)
	}
}
