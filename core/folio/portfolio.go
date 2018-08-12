package folio

import "github.com/louisevanderlith/husk"

type Portfolio struct {
	ImageID int64    `hsk:"null"`
	URL     string   `hsk:"size(128)"`
	Name    string   `hsk:"size(50)"`
	Profile *Profile `json:",omitempty"`
}

func (o Portfolio) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func getPortfolio(id int64) (portfolioRecord, error) {
	return ctx.Portfolios.FindByID(id)
}

func GetPortfolio(id int64) (*Portfolio, error) {
	rec, err := getPortfolio(id)

	return rec.Data(), err
}

func GetPortfolios(page, size int) (portfolioSet, error) {
	return ctx.Portfolios.Find(page, size, func(o Portfolio) bool {
		return true
	})
}

func (p Portfolio) Update(id int64) error {
	prtfolio, err := getPortfolio(id)

	if err != nil {
		return err
	}

	prtfolio.Set()

	return ctx.Portfolios.Update(prtfolio)
}

func (p *Portfolio) AddSocialLink(id int64, socialLink SocialLink) error {
	profile := p.Profile
	profile.SocialLinks = append(profile.SocialLinks, socialLink)

	return p.Update(id)
}
