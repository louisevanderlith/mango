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
	defer createDefaultWebsite()

	ctx = context{
		Profiles: husk.NewTable(new(Profile)),
	}
}

func createDefaultWebsite() {
	if ctx.Profiles.Exists(husk.Everything()) {
		return
	}

	vosa := Profile{
		Title:        "avosa",
		ContactEmail: "abc@vosa.com",
		Description:  "software for the automotive industry",
		ContactPhone: "0893523423",
		URL:          "https://www.localhost.co.za",
	}

	rec := vosa.Create()
	log.Printf("Default: %#v\n", rec)
	if rec.Error != nil {
		panic(rec.Error)
	}

	defer ctx.Profiles.Save()
	log.Printf("Default Website Loaded:\n%+v\n", rec.Record.Data())
}
