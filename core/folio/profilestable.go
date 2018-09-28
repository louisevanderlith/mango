package folio

import (
	"log"

	"github.com/louisevanderlith/husk"
)

type profilesTable struct {
	tbl husk.Tabler
}

func NewProfilesTable() profilesTable {
	result := husk.NewTable(new(Profile))

	return profilesTable{result}
}

func (t profilesTable) FindByKey(key husk.Key) (profileRecord, error) {
	result, err := t.tbl.FindByKey(key)

	return profileRecord{result}, err
}

func (t profilesTable) Find(page, pageSize int, filter profileFilter) (profileSet, error) {
	huskFilter, err := husk.MakeFilter(filter)

	var result profileSet

	if err == nil {
		items := t.tbl.Find(page, pageSize, huskFilter)
		result = profileSet{items}
	}

	return result, err
}

func (t profilesTable) FindFirst(filter profileFilter) (profileRecord, error) {
	huskFilter, err := husk.MakeFilter(filter)

	if err != nil {
		return profileRecord{}, err
	}

	result := t.tbl.FindFirst(huskFilter)

	return profileRecord{result}, err
}

func (t profilesTable) Exists(filter profileFilter) bool {
	huskFilter, err := husk.MakeFilter(filter)

	if err != nil {
		log.Println(err)
		return true
	}

	return t.tbl.Exists(huskFilter)
}

func (t profilesTable) Create(obj Profile) (profileRecord, error) {
	set := t.tbl.Create(obj)
	defer t.tbl.Save()

	return profileRecord{set.Record}, set.Error
}

func (t profilesTable) Update(record profileRecord) error {
	result := t.tbl.Update(record.rec)
	defer t.tbl.Save()

	return result
}

func (t profilesTable) Delete(key husk.Key) error {
	return t.tbl.Delete(key)
}

type profileRecord struct {
	rec husk.Recorder
}

func (r profileRecord) Data() *Profile {
	return r.rec.Data().(*Profile)
}

func (r *profileRecord) Set(profile Profile) error {
	return r.rec.Set(profile)
}

type profileFilter func(o Profile) bool

type profileSet struct {
	*husk.RecordSet
}

func newprofileSet() *profileSet {
	result := husk.NewRecordSet()

	return &profileSet{result}
}
