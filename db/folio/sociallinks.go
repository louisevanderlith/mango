package folio

import "github.com/louisevanderlith/db"

type SocialLinks []*SocialLink

func (o SocialLinks) Each(handler func(db.IRecord)) {
	for _, v := range o {
		handler(v)
	}
}

func (o SocialLinks) Length() int {
	return len(o)
}

func (o SocialLinks) At(index int) db.IRecord {
	return o[index]
}

func (o SocialLinks) Add(obj db.IRecord) {
	item, ok := obj.(*SocialLink)

	if ok {
		o = append(o, item)
	}
}
