package folio

type context struct {
	Abouts      aboutsTable
	Portfolios  portfoliosTable
	Profiles    profilesTable
	SocialLinks socialLinksTable
	Headers     headersTable
}

var ctx context

func NewContext() {
	ctx = context{
		Abouts:      NewAboutsTable(),
		Portfolios:  NewPortfoliosTable(),
		Profiles:    NewProfilesTable(),
		SocialLinks: NewSocialLinksTable(),
		Headers:     NewHeadersTable(),
	}
}
