package classifieds

import (
	"strings"
	"time"

	"errors"

	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/husk"
)

type CarAdvert struct {
	db.Record
	ModelID       int64
	Info          string `hsk:"size(128)"`
	Year          int    `orm:"null"`
	Odometer      int    `orm:"null"`
	HasPapers     bool   `hsk:"default(false)"`
	LicenseExpiry time.Time
}

func (o CarAdvert) Valid() (bool, error) {
	var issues []string

	valid, common := husk.ValidateStruct(&o)
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
