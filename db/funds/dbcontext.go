package funds

type context struct {
	Experiences  experiencesTable
	Heroes       heroesTable
	Levels       levelsTable
	LineItems    lineItemsTable
	Requisitions requisitionsTable
	Transactions transactionsTable
}

var ctx context

func NewContext() {
	ctx = context{
		Experiences:  NewExperiencesTable(),
		Heroes:       NewHeroesTable(),
		Levels:       NewLevelsTable(),
		LineItems:    NewLineItemsTable(),
		Requisitions: NewRequisitionsTable(),
		Transactions: NewTransactionsTable(),
	}

	go seedData()
}

func seedData() {
	seedLevel()
}
