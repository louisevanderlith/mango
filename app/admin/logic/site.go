package logic

import (
	"github.com/louisevanderlith/mango/util"
	"encoding/json"
	"log"
	"errors"
)

type BasicSite struct {
	ID          int64
	Title       string
	Description string
}

func GetSites() (result []BasicSite, finalErr error) {
	contents, statusCode := util.GETMessage("Folio.API", "site")
	data := util.MarshalToMap(contents)

	if statusCode != 200 {
		var dataErr string
		err := json.Unmarshal(*data["Error"], &dataErr)

		if err != nil {
			log.Printf("getSites: ", err)
		}

		finalErr = errors.New(dataErr)
	} else {
		err := json.Unmarshal(*data["Data"], &result)

		if err != nil {
			log.Printf("getSites: ", err)
		}
	}

	return result, finalErr
}
