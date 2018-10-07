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

func getProfile(key *husk.Key) (husk.Recorder, error) {
	return ctx.Profiles.FindByKey(key)
}

func getProfileByName(name string) husk.Recorder {
	return ctx.Profiles.FindFirst(byName(name))
}

func GetProfile(key *husk.Key) (*Profile, error) {
	rec, err := getProfile(key)

	if err != nil {
		return nil, err
	}

	return rec.Data().(*Profile), nil
}

func GetProfileByName(name string) *Profile {
	rec := getProfileByName(name)

	return rec.Data().(*Profile)
}

func GetProfiles(page, size int) husk.Collection {
	return ctx.Profiles.Find(page, size, husk.Everything())
}

func (p Profile) Create() husk.CreateSet {
	return ctx.Profiles.Create(p)
}

func (p Profile) Update(key *husk.Key) error {
	profile, err := getProfile(key)

	if err != nil {
		return err
	}

	profile.Set(p)

	return ctx.Profiles.Update(profile)
}

func AddSocialLink(key *husk.Key, socialLink SocialLink) error {
	prRec, err := getProfile(key)

	if err != nil {
		return err
	}

	profile := prRec.Data().(*Profile)
	profile.SocialLinks = append(profile.SocialLinks, socialLink)

	return ctx.Profiles.Update(prRec)
}

func AddAboutSection(key *husk.Key, about About) error {
	prRec, err := getProfile(key)

	if err != nil {
		return err
	}

	profile := prRec.Data().(*Profile)
	profile.AboutSections = append(profile.AboutSections, about)

	return ctx.Profiles.Update(prRec)
}

func AddHeaderSection(key *husk.Key, header Header) error {
	prRec, err := getProfile(key)

	if err != nil {
		return err
	}

	profile := prRec.Data().(*Profile)
	profile.Headers = append(profile.Headers, header)

	return ctx.Profiles.Update(prRec)
}

func AddPortfolioSection(key *husk.Key, portfolio Portfolio) error {
	prRec, err := getProfile(key)

	if err != nil {
		return err
	}

	profile := prRec.Data().(*Profile)
	profile.PortfolioItems = append(profile.PortfolioItems, portfolio)

	return ctx.Profiles.Update(prRec)
}
