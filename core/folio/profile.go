package folio

import "github.com/louisevanderlith/husk"

type Profile struct {
	Title          string       `hsk:"size(128)" json:",omitempty"`
	Description    string       `hsk:"size(512)" json:",omitempty"`
	ContactEmail   string       `hsk:"size(128)" json:",omitempty"`
	ContactPhone   string       `hsk:"size(20)" json:",omitempty"`
	URL            string       `hsk:"size(128)" json:",omitempty"`
	ImageID        int64        `hsk:"null"`
	StyleSheet     string       `hsk:"size(50)"`
	SocialLinks    []SocialLink `json:",omitempty"`
	PortfolioItems []Portfolio  `json:",omitempty"`
	AboutSections  []About      `json:",omitempty"`
	Headers        []Header     `json:",omitempty"`
}

func (p Profile) Valid() (bool, error) {
	return husk.ValidateStruct(&p)
}
