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
	ID         int64     `orm:"column(id);auto;pk"`
	CreateDate time.Time `orm:"auto_now_add;type(time)"`
	Deleted    bool      `orm:"default(false)"`
}

func (r Record) GetID() int64 {
	return r.ID
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