package classifieds

import (
	"time"

	"github.com/louisevanderlith/mango/db"
	"errors"
	"strings"
	"github.com/louisevanderlith/mango/util"
)

type CarAdvert struct {
	db.Record
	ModelID         int64
	Info          string    `orm:"size(128)"`
	Year          int       `orm:"null"`
	Odometer      int       `orm:"null"`
	HasPapers     bool      `orm:"default(false)"`
	LicenseExpiry time.Time `orm:"type(date)"`
}

func (o CarAdvert) Validate() (bool, error) {
	var issues []string

	valid, common := util.ValidateStruct(o)
	if !valid {
		issues = append(issues, common.Error())
	}

	if o.Year > 0 && o.Year > time.Now().Year() {
		issues = append(issues, "Year can't be in the future.")
	}

	if o.Odometer < 0 {
		issues = append(issues, "Odometer can't be negative.")
	}

	if o.HasPapers && o.LicenseExpiry.Before(time.Now()) {
		issues = append(issues, "License has already expired.")
	}

	isValid := len(issues) < 1
	finErr := errors.New(strings.Join(issues, "\r\n"))

	return isValid, finErr
}