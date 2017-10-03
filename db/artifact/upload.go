package artifact

type Upload struct {
	Record
	ItemID   int64
	Name     string
	MimeType string
	Path     string
	Size     int64
}
