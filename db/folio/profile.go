package folio

import "github.com/louisevanderlith/husk"

type Profile struct {
	Title          string      `hsk:"size(128)" json:",omitempty"`
	Description    string      `hsk:"size(512)" json:",omitempty"`
	ContactEmail   string      `hsk:"size(128)" json:",omitempty"`
	ContactPhone   string      `hsk:"size(20)" json:",omitempty"`
	URL            string      `hsk:"size(128)" json:",omitempty"`
	ImageID        int64       `hsk:"null"`
	StyleSheet     string      `hsk:"size(50)"`
	SocialLinks    SocialLinks `json:",omitempty"`
	PortfolioItems Portfolios  `json:",omitempty"`
	AboutSections  Abouts      `json:",omitempty"`
	Headers        Headers     `json:",omitempty"`
}

func (p Profile) Valid() (bool, error) {
	return husk.ValidateStruct(&p)
}
