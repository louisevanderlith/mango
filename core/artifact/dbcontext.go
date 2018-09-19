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

func Shutdown() {
	ctx.Uploads.Save()
}
