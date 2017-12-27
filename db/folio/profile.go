package folio

import (
	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
)

type Profile struct {
	db.Record
	Title          string        `orm:"size(128)"`
	Description    string        `orm:"size(512)"`
	ContactEmail   string        `orm:"size(128)"`
	ContactPhone   string        `orm:"size(20)"`
	URL            string        `orm:"size(128)"`
	ImageURL       string        `orm:"size(85)"`
	StyleSheet     string        `orm:"size(50)"`
	SocialLinks    []*SocialLink `orm:"reverse(many)"`
	PortfolioItems []*Portfolio  `orm:"reverse(many)"`
	AboutSections  []*About      `orm:"reverse(many)"`
}

func (p Profile) Validate() (bool, error) {
	return util.ValidateStruct(p)
}
