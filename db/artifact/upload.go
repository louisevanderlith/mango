package artifact

import "github.com/louisevanderlith/mango/util"

type Upload struct {
	util.Record
	ItemID   int64
	Name     string
	MimeType string
	Path     string
	Size     int64
}
