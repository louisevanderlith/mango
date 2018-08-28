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

func getProfile(key husk.Key) (profileRecord, error) {
	return ctx.Profiles.FindByKey(key)
}

func GetProfile(key husk.Key) (*Profile, error) {
	rec, err := getProfile(key)

	return rec.Data(), err
}

func GetProfiles(page, size int) (profileSet, error) {
	return ctx.Profiles.Find(page, size, func(o Profile) bool {
		return true
	})
}

func (p Profile) Update(key husk.Key) error {
	prtfolio, err := getProfile(key)

	if err != nil {
		return err
	}

	prtfolio.Set(p)

	return ctx.Profiles.Update(prtfolio)
}

func (p *Profile) AddSocialLink(key husk.Key, socialLink SocialLink) error {
	p.SocialLinks = append(p.SocialLinks, socialLink)

	return p.Update(key)
}
