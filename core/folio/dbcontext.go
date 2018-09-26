package folio

import (
	"log"
)

type context struct {
	Profiles profilesTable
}

var ctx context

func init() {
	ctx = context{
		Profiles: NewProfilesTable(),
	}

	createDefaultWebsite()
}

func createDefaultWebsite() {
	any, _ := ctx.Profiles.Exists(func(obj Profile) bool {
		return true
	})

	if any {
		return
	}

	vosa := Profile{
		Title:        "avosa",
		ContactEmail: "abc@vosa.com",
		Description:  "software for the automotive industry",
		ContactPhone: "0893523423",
		URL:          "https://www.localhost.co.za",
		StyleSheet:   "avosa.css",
	}

	rec, err := vosa.Create()

	if err != nil {
		panic(err)
	}

	log.Printf("Default Website Loaded:\n%+v\n", rec)
}
