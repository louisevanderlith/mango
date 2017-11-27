package artifact

import (
	"github.com/louisevanderlith/mango/db"
)

type Upload struct {
	db.Record
	ItemID   int64
	Name     string
	MimeType string
	Path     string
	Size     int64
}
