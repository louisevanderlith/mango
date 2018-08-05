package artifact

type context struct {
	Uploads uploadsTable
	BLOBS   []Blob
}

var ctx context

func NewContext() {
	ctx = context{
		Uploads: NewUploadsTable(),
	}
}
