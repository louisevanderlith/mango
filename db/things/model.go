package things

type Model struct {
	db.Record
	ManufacturerID *Manufacturer
	Name string
}