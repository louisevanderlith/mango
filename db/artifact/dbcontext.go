package artifact

type context struct {
	Uploads uploadsTable
	BLOBS   blobsTable
}

var ctx context

func NewContext() {
	ctx = context{
		Uploads: NewUploadsTable(),
		BLOBS:   NewBLOBSTable(),
	}
}
