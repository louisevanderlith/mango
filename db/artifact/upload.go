package artifact

import "github.com/louisevanderlith/mango/util"

type Upload struct {
	util.BaseRecord
	ItemID   int64
	Name     string
	MimeType string
	Path     string
	Size     int64
}
