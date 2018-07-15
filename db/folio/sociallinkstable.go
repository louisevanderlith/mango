package folio

import (
	"github.com/louisevanderlith/husk"
)

type socialLinksTable struct {
	tbl husk.Tabler
}

func NewSocialLinksTable() socialLinksTable {
	result := husk.NewTable(new(SocialLink))

	return socialLinksTable{result}
}

func (t socialLinksTable) Create(obj SocialLink) (socialLinkRecord, error) {
	result, err := t.tbl.Create(obj)

	return socialLinkRecord{result}, err
}

type socialLinkRecord struct {
	rec husk.Recorder
}
