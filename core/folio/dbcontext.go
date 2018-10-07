package folio

import (
	"log"

	"github.com/louisevanderlith/husk"
)

type context struct {
	Profiles husk.Tabler
}

var ctx context

func init() {
	ctx = context{
		Profiles: husk.NewTable(new(Profile)),
	}

	createDefaultWebsite()
}

func createDefaultWebsite() {
	any := ctx.Profiles.Exists(husk.Everything())

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

	rec := vosa.Create()

	if rec.Error != nil {
		panic(rec.Error)
	}

	log.Printf("Default Website Loaded:\n%+v\n", rec)
}
