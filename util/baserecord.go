package util

import (
	"time"
)

type BaseRecord struct {
	ID         int64     `orm:"column(id);auto;pk"`
	CreateDate time.Time `orm:"auto_now_add;type(datetime)"`
	Deleted    bool      `orm:"default(false)"`
}
