package game

import (
	"time"

	"github.com/louisevanderlith/husk"
)

type Hero struct {
	UserID      int64
	Credits     int
	Experiences []Experience
	Level       *Level
	TotalXP     int
	LastUpdated time.Time //update on save???
}

func (o Hero) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func (h *Hero) AddExperience(xpType ExperienceType) {
	// Add XP Record with (new) Level
	// Update Hero to reflect new TotalXP
	xpValue := XPValue(xpType)
	xp := Experience{
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
