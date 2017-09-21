package things

type Category struct {
	db.Record
	Name string
	Description string
}