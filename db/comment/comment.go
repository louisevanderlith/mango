package comment

type Comment struct {
	Record
	User      *User `orm:"null;rel(one)"`
	UpVotes   int
	DownVotes int
	Adverts   []*Advert `orm:"reverse(many)"`
	Profile   *Profile  `orm:"rel(fk)"`
}
