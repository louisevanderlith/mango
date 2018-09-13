package logic

import (
	"errors"
	"log"
	"strconv"

	"github.com/louisevanderlith/mango/util"
)

type BasicSite struct {
	ID             int64
	Title          string
	Description    string
	ContactEmail   string
	ContactPhone   string
	URL            string
	ImageID        int64
	ImageURL       string
	StyleSheet     string
	SocialLinks    SocialLinks
	PortfolioItems PortfolioItems
	AboutSections  AboutSections
	Headers        HeaderItems
}

type SocialLinks []socialLink

type PortfolioItems []portfolioItem

type AboutSections []aboutSection

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

func GetSites() (result []BasicSite, finalErr error) {
	contents, err := util.GETMessage("Folio.API", "site")

	if err != nil {
		return result, err
	}

	data := util.MarshalToResult(contents)

	if len(data.Error) != 0 {
		return result, errors.New(data.Error)
	}

	result = data.Data.([]BasicSite)

	return result, nil
}

func GetSite(siteID int64) (result BasicSite, finalErr error) {
	contents, err := util.GETMessage("Folio.API", "site", strconv.FormatInt(siteID, 10))

	if err != nil {
		return result, err
	}

	data := util.MarshalToResult(contents)

	if len(data.Error) != 0 {
		return result, errors.New(data.Error)
	}

	result = data.Data.(BasicSite)
	result.setImageURLs()

	return result, nil
}

func (obj *BasicSite) setImageURLs() {
	if uploadURL == "" {
		setUploadURL()
	}

	if obj.ImageID != 0 {
		obj.ImageURL = uploadURL + strconv.FormatInt(obj.ImageID, 10)
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

func setUploadURL() {
	url, err := util.GetServiceURL("Artifact.API", true)

	if err != nil {
		log.Print("setImageURLs:", err)
	}

	uploadURL = url + "v1/upload/file/"
}
