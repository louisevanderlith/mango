package comment

type Comment struct {
	Record
	UserID    int64 `orm:"null"`
	UpVotes   int
	DownVotes int
	ItemID    int64 `orm:"null"`
}
