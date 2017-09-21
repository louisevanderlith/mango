package classifieds

type Tag struct {
	Record
	Description string
	Adverts      []*Advert `orm:"reverse(many)"`
}
