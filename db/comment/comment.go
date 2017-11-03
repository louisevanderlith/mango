package comment

import "github.com/louisevanderlith/mango/util"

type Comment struct {
	util.BaseRecord
	UserID    int64 `orm:"null"`
	UpVotes   int
	DownVotes int
	ItemID    int64 `orm:"null"`
}

func (obj *Comment) Insert() (int64, error) {
	return db.Insert(obj)
}

func (obj *Comment) Read() error {
	return db.Read(*obj)
}

func (obj *Comment) Update() (int64, error) {
	return db.Update(obj)
}

func (obj *Comment) Delete() error {
	_, err := db.Delete(obj)

	return err
}