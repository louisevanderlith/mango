package auto

type context struct {
	Adverts advertsTable
}

var ctx context

func init() {
	ctx = context{
		Adverts: NewAdvertsTable(),
	}
}
