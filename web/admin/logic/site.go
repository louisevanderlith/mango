package logic

import (
	"github.com/louisevanderlith/mango/core/folio"

	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/mango/pkg"
)

var uploadURL string

func GetSites(instanceID string) ([]folio.Profile, error) {
	var result []folio.Profile

	resp, err := util.GETMessage(instanceID, "Folio.API", "profile", "all", "A10")

	if err != nil {
		return result, err
	}

	if resp.Failed() {
		return result, resp
	}

	coll, ok := resp.Data.(husk.Collection)

	if ok && coll.Any() {
		itor := coll.GetEnumerator()

		if itor.MoveNext() {
			curr := itor.Current()

			result = append(result, curr.Data().(folio.Profile))
		}
	}

	return result, nil
}

func GetSite(siteKey *husk.Key, instanceID string) (folio.Profile, error) {
	result := folio.Profile{}
	resp, err := util.GETMessage(instanceID, "Folio.API", "profile", siteKey.String())

	if err != nil {
		return result, err
	}

	if resp.Failed() {
		return result, resp
	}

	result = resp.Data.(folio.Profile)
	//result.setImageURLs(instanceID)

	return result, nil
}

/*
func setImageURLs(instanceID string) {
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
*/
