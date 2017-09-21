package artifact

type Upload struct{
	Record
	Advert *Advert `orm:"rel(fk)"`
}