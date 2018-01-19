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
}

type SocialLinks []socialLink

type PortfolioItems []portfolioItem

type AboutSections []aboutSection

type socialLink struct {
	ID   int64
	Icon string
	URL  string
}

type portfolioItem struct {
	ImageID  int64
	ImageURL string
	URL      string
	Name     string
}

type aboutSection struct {
	ID          int64
	SectionText string
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
			log.Print("getSite: ", err)
		}

		finalErr = errors.New(dataErr)
	} else {
		err := json.Unmarshal(*data["Data"], &result)

		if err != nil {
			log.Print("getSite: ", err)
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

	for _, v := range obj.PortfolioItems {
		v.ImageURL = uploadURL + strconv.FormatInt(v.ImageID, 10)
	}
}

func setUploadURL() {
	url, err := util.GetServiceURL("Artifact.API", true)

	if err != nil {
		log.Print("setImageURLs:", err)
	}

	uploadURL = url + "v1/upload/file/"
}
