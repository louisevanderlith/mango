package folio

type context struct {
	Portfolios portfoliosTable
}

var ctx context

func init() {
	ctx = context{
		Portfolios: NewPortfoliosTable(),
	}
}
