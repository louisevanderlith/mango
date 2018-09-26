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

func getProfileByName(name string) (profileRecord, error) {
	return ctx.Profiles.FindFirst(func(obj Profile) bool {
		return obj.Title == name
	})
}

func GetProfile(key husk.Key) (*Profile, error) {
	rec, err := getProfile(key)

	if err != nil {
		return nil, err
	}

	return rec.Data(), nil
}

func GetProfileByName(name string) (*Profile, error) {
	rec, err := getProfileByName(name)

	if err != nil {
		return nil, err
	}

	return rec.Data(), nil
}

func GetProfiles(page, size int) (profileSet, error) {
	return ctx.Profiles.Find(page, size, func(o Profile) bool {
		return true
	})
}

func (p Profile) Create() (profileRecord, error) {
	return ctx.Profiles.Create(p)
}

func (p Profile) Update(key husk.Key) error {
	profile, err := getProfile(key)

	if err != nil {
		return err
	}

	profile.Set(p)

	return ctx.Profiles.Update(profile)
}

func AddSocialLink(key husk.Key, socialLink SocialLink) error {
	prRec, err := getProfile(key)

	if err != nil {
		return err
	}

	profile := prRec.Data()
	profile.SocialLinks = append(profile.SocialLinks, socialLink)

	return ctx.Profiles.Update(prRec)
}

func AddAboutSection(key husk.Key, about About) error {
	prRec, err := getProfile(key)

	if err != nil {
		return err
	}

	profile := prRec.Data()
	profile.AboutSections = append(profile.AboutSections, about)

	return ctx.Profiles.Update(prRec)
}

func AddHeaderSection(key husk.Key, header Header) error {
	prRec, err := getProfile(key)

	if err != nil {
		return err
	}

	profile := prRec.Data()
	profile.Headers = append(profile.Headers, header)

	return ctx.Profiles.Update(prRec)
}

func AddPortfolioSection(key husk.Key, portfolio Portfolio) error {
	prRec, err := getProfile(key)

	if err != nil {
		return err
	}

	profile := prRec.Data()
	profile.PortfolioItems = append(profile.PortfolioItems, portfolio)

	return ctx.Profiles.Update(prRec)
}
