package secure

type Profile struct {
	Record
	Name      string `orm:"size(75)"`
	Verified  bool
	UpVotes   int
	DownVotes int
	Comments  []*Comment `orm:"reverse(many)"`
	User      *User      `orm:"reverse(one)"`
}
