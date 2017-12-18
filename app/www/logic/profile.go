package logic

import (
	"github.com/louisevanderlith/mango/db/folio"
	"github.com/louisevanderlith/mango/util"
	"encoding/json"
	"log"
	"errors"
)

func GetProfileSite(name string) (result folio.Profile, finalErr error) {
	if name == "" {
		name = "avosa"
	}

	contents, statusCode := util.GETMessage("Folio.API", "site", name)
	data := util.MarshalToMap(contents)

	if statusCode != 200 {
		var dataErr string
		err := json.Unmarshal(*data["Error"], &dataErr)

		if err != nil {
			log.Printf("getProfileSite: ", err)
		}

		finalErr = errors.New(dataErr)
	} else {
		finalErr = json.Unmarshal(*data["Data"], &result)
	}

	return result, finalErr
}
