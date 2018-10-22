package logic

import (
	"log"
	"strconv"

	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/mango/util"
)

type BasicSite struct {
	Key            husk.Key
	Title          string
	Description    string
	ContactEmail   string
	ContactPhone   string
	URL            string
	SocialLinks    SocialLinks
	PortfolioItems PortfolioItems
	AboutSections  []string
	Headers        HeaderItems
}

type SocialLinks []socialLink

type PortfolioItems []portfolioItem

type HeaderItems []headerItem

type socialLink struct {
	ID   int64
	Icon string
	URL  string
}

type portfolioItem struct {
	ID       int64
	ImageID  int64
	ImageURL string
	URL      string
	Name     string
}

type aboutSection struct {
	ID          int64
	SectionText string
}

type headerItem struct {
	ID       int64
	Heading  string
	Text     string
	ImageID  int64
	ImageURL string
}

var uploadURL string

func GetSites(instanceID string) ([]BasicSite, error) {
	var result []BasicSite

	resp, err := util.GETMessage(instanceID, "Folio.API", "profile")

	if err != nil {
		return result, err
	}

	if resp.Failed() {
		return result, resp
	}

	result = resp.Data.([]BasicSite)

	return result, nil
}

func GetSite(siteKey *husk.Key, instanceID string) (BasicSite, error) {
	result := BasicSite{}
	resp, err := util.GETMessage(instanceID, "Folio.API", "profile", siteKey.String())

	if err != nil {
		return result, err
	}

	if resp.Failed() {
		return result, resp
	}

	result = resp.Data.(BasicSite)
	result.setImageURLs(instanceID)

	return result, nil
}

func (obj *BasicSite) setImageURLs(instanceID string) {
	if uploadURL == "" {
		setUploadURL(instanceID)
	}

	for i := 0; i < len(obj.PortfolioItems); i++ {
		row := &obj.PortfolioItems[i]

		if row.ImageID != 0 {
			row.ImageURL = uploadURL + strconv.FormatInt(row.ImageID, 10)
		}
	}

	for i := 0; i < len(obj.Headers); i++ {
		row := &obj.Headers[i]

		if row.ImageID != 0 {
			row.ImageURL = uploadURL + strconv.FormatInt(row.ImageID, 10)
		}
	}
}

func setUploadURL(instanceID string) {
	url, err := util.GetServiceURL(instanceID, "Artifact.API", true)

	if err != nil {
		log.Print("setImageURLs:", err)
	}

	uploadURL = url + "v1/upload/file/"
}
