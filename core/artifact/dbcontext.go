package artifact

type context struct {
	Uploads uploadsTable
}

var ctx context

func init() {
	ctx = context{
		Uploads: NewUploadsTable(),
	}
}
