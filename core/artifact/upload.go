package artifact

import (
	"github.com/louisevanderlith/husk"
)

type Upload struct {
	ItemID   int64
	ItemName string `hsk:"size(75)"`
	Name     string `hsk:"size(50)"`
	MimeType string `hsk:"size(30)"`
	Size     int64
	BLOB     *Blob
}

func (o Upload) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func GetUploads(page, pagesize int) husk.Collection {
	return ctx.Uploads.Find(page, pagesize, husk.Everything())
}

func GetUpload(key *husk.Key) (result husk.Recorder, err error) {
	return ctx.Uploads.FindByKey(key)
}

func GetUploadFile(key *husk.Key) (result []byte, filename string, err error) {
	upload, err := GetUpload(key)

	if err != nil {
		return nil, "", err
	}

	uploadData := upload.Data().(*Upload)
	blob := uploadData.BLOB.Data

	return blob, uploadData.Name, err
}

//GetUploadsBySize returns the first 50 records larger than @size bytes.
func GetUploadsBySize(size int64) husk.Collection {
	return ctx.Uploads.Find(1, 50, bySize(size))
}

func (upload Upload) Create() husk.CreateSet {
	return ctx.Uploads.Create(upload)
}
