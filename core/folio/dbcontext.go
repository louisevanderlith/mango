package folio

type context struct {
	Profiles profilesTable
}

var ctx context

func init() {
	ctx = context{
		Profiles: NewProfilesTable(),
	}
}
