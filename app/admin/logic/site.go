package logic

import (
	"encoding/json"
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
	Text     string
	ImageID  int64
	ImageURL string
}

var uploadURL string

func GetSites() (result []BasicSite, finalErr error) {
	contents, statusCode := util.GETMessage("Folio.API", "site")
	data := util.MarshalToMap(contents)

	if statusCode != 200 {
		var dataErr string
		err := json.Unmarshal(*data["Error"], &dataErr)

		if err != nil {
			log.Print("getSites: ", err)
		}

		finalErr = errors.New(dataErr)
	} else {
		err := json.Unmarshal(*data["Data"], &result)

		if err != nil {
			log.Print("getSites: ", err)
		}
	}

	return result, finalErr
}

func GetSite(siteID int64) (result BasicSite, finalErr error) {
	contents, statusCode := util.GETMessage("Folio.API", "site", strconv.FormatInt(siteID, 10))
	data := util.MarshalToMap(contents)

	if statusCode != 200 {
		var dataErr string
		err := json.Unmarshal(*data["Error"], &dataErr)

		if err != nil {
			log.Println("getSite: ", err)
		}

		finalErr = errors.New(dataErr)
	} else {
		err := json.Unmarshal(*data["Data"], &result)

		if err != nil {
			log.Println("getSite: ", err)
		}

		result.setImageURLs()
	}

	return result, finalErr
}

func (obj *BasicSite) setImageURLs() {
	if uploadURL == "" {
		setUploadURL()
	}

	obj.ImageURL = uploadURL + strconv.FormatInt(obj.ImageID, 10)

	for i := 0; i < len(obj.PortfolioItems); i++ {
		row := &obj.PortfolioItems[i]
		row.ImageURL = uploadURL + strconv.FormatInt(row.ImageID, 10)
	}

	for i := 0; i < len(obj.HeaderItems); i++ {
		row := &obj.HeaderItems[i]
		row.ImageURL = uploadURL + strconv.FormatInt(row.ImageID, 10)
	}
}

func setUploadURL() {
	url, err := util.GetServiceURL("Artifact.API", true)

	if err != nil {
		log.Print("setImageURLs:", err)
	}

	uploadURL = url + "v1/upload/file/"
}
