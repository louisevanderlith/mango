package artifact

type Upload struct{
	db.Record
	Advert *Advert `orm:"rel(fk)"`
}