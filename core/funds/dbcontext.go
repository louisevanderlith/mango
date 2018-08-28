package funds

type context struct {
	Transactions transactionsTable
}

var ctx context

func init() {
	ctx = context{
		Transactions: NewTransactionsTable(),
	}
}
