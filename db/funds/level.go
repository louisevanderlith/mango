package funds

import (
	"github.com/louisevanderlith/husk"
)

type Level struct {
	Rank     int
	Required int
	Next     *Level
}

func (o Level) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func getLevel(xp int) husk.Recorder {
	// return level based on xp given
	return ctx.Levels.FindFirst(func(o husk.Dataer) bool {
		item := o.(*Level)
		return xp <= item.Required
	})
}

func seedLevel() {
	exists := ctx.Levels.Exists(func(o husk.Dataer) bool {
		return true
	})

	if !exists {
		for i := 75; i >= 0; i-- {
			ctx.Levels.Create(createLevel(i))
		}
	}
}

func createLevel(lvl int) Level {
	result := Level{}

	result.Rank = lvl
	result.Required = xpRequired(lvl)

	if lvl != 75 {
		nxt := createLevel(lvl + 1)
		result.Next = &nxt
	}

	return result
}

func xpRequired(lvl int) int {
	if lvl == 0 || lvl == 1 {
		return lvl * 50.0
	}

	return xpRequired(lvl-1) + lvl*50
}
