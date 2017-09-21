package classifieds

import "time"

type CarAdvert struct {
	Record
	Make          string    `orm:"size(50)"`
	Model         string    `orm:"size(50)"`
	Info          string    `orm:"size(128)"`
	Year          int       `orm:"null"`
	Odometer      int       `orm:"null"`
	HasPapers     bool      `orm:"default(false)"`
	LicenseExpiry time.Time `orm:"type(date)"`
}
