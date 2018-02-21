package db

import "time"

type IRecord interface {
	GetID() int64
	IsDeleted() bool
	GetCreateDate() time.Time
	Disable() Record
	Validate() (bool, error)
	Exists() (bool, error)
}

type Record struct {
	Id         int64     `orm:"auto;pk;unique;`
	CreateDate time.Time `orm:"auto_now_add"`
	Deleted    bool      `orm:"default(false)"`
}

func (r Record) GetID() int64 {
	return r.Id
}

func (r Record) IsDeleted() bool {
	return r.Deleted
}

func (r Record) GetCreateDate() time.Time {
	return r.CreateDate
}

func (r Record) Disable() Record {
	r.Deleted = true

	return r
}

func (r Record) Exists() (bool, error) {
	return false, nil
}
