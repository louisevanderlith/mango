package game

type context struct {
	Heroes heroesTable
}

var ctx context

func init() {
	ctx = context{
		Heroes: NewHeroesTable(),
	}

	go seedLevel()
}
