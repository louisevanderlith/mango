package classifieds

type context struct {
	Adverts    advertsTable
	CarAdverts carAdvertsTable
	Tags       tagsTable
}

var ctx context

func NewContext() {
	ctx = context{
		Adverts:    NewAdvertsTable(),
		CarAdverts: NewCarAdvertsTable(),
		Tags:       NewTagsTable(),
	}
}
