package funds

import (
	"fmt"

	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Level struct {
	db.Record
	Rank     int
	Required int
	Next     *Level
}

func (o Level) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}

func getLevel(xp int) *Level {
	// return level based on xp given
	result := new(Level)

	var container Levels
	err := Ctx.Levels.Read(&Level{}, &container)

	if err == nil {
		for _, v := range container {
			if xp <= v.Required {
				result = v
				break
			}
		}
	} else {
		fmt.Print(err)
	}

	return result
}

func seedLevel() {
	var items Levels
	err := Ctx.Levels.Read(&Level{}, &items)

	if err == nil {
		if len(items) == 0 {
			var data Levels
			for i := 75; i >= 0; i-- {
				data.Add(createLevel(i))
			}

			_, err = Ctx.Levels.CreateMulti(data)

			if err != nil {
				fmt.Println("seedLevel:", err)
			}
		}
	} else {
		fmt.Println("seedLevel:", err)
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
