package logic

import (
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

func GetProfileSite(instanceID, name string) (result BasicSite, finalErr error) {
	if name == "" {
		name = "avosa"
	}

	contents, err := util.GETMessage(instanceID, "Folio.API", "site", name)

	if err != nil {
		return result, err
	}

	data := util.MarshalToResult(contents)

	if data.Failed() {
		return result, data
	}

	result = data.Data.(BasicSite)
	result.setImageURLs(instanceID)

	return result, nil
}

func (obj *BasicSite) setImageURLs(instanceID string) {
	if uploadURL == "" {
		setUploadURL(instanceID)
	}

	obj.ImageURL = uploadURL + strconv.FormatInt(obj.ImageID, 10)

	for i := 0; i < len(obj.PortfolioItems); i++ {
		row := &obj.PortfolioItems[i]
		row.ImageURL = uploadURL + strconv.FormatInt(row.ImageID, 10)
	}

	for i := 0; i < len(obj.Headers); i++ {
		row := &obj.Headers[i]
		row.ImageURL = uploadURL + strconv.FormatInt(row.ImageID, 10)
	}
}

func setUploadURL(instanceID string) {
	url, err := util.GetServiceURL(instanceID, "Artifact.API", true)

	if err != nil {
		log.Print("setImageURLs:", err)
	}

	uploadURL = url + "v1/upload/file/"
}
