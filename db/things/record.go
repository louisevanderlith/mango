package things

import (
	"time"
)

type Record struct {
	ID         int64     `orm:"column(Id);pk"`
	CreateDate time.Time `orm:"auto_now_add;type(date)"`
	Deleted    bool      `orm:"default(false)"`
}
