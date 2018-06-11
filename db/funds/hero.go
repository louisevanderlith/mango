package funds

import (
	"time"

	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Hero struct {
	db.Record
	UserID       int64
	Credits      int
	Requisitions Requisitions
	Experiences  Experiences
	Level        *Level
	TotalXP      int
	LastUpdated  time.Time `orm:"auto_update_???`
}

func (o Hero) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}

func (h *Hero) AddExperience(xpType ExperienceType) {
	// Add XP Record with (new) Level
	// Update Hero to reflect new TotalXP
	xpValue := XPValue(xpType)
	xp := Experience{
		Hero:   h,
		Points: xpValue,
		Type:   xpType,
	}

	h.Experiences = append(h.Experiences, &xp)
	h.TotalXP += xpValue

}

func (h *Hero) AddRequisition() {

}

func (h *Hero) AddCredit(amount int) {

}
