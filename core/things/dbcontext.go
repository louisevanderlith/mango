package things

type context struct {
	Categories    categoriesTable
	Manufacturers manufacturersTable
	Models        modelsTable
	SubCategories subcategoriesTable
}

var ctx context

func NewContext() {
	ctx = context{
		Categories:    NewCategoriesTable(),
		Manufacturers: NewManufacturersTable(),
		Models:        NewModelsTable(),
		SubCategories: NewSubcategoriesTable(),
	}

	go seedData()
}

func seedData() {
	seedManufacturer()
	seedModel()
}
