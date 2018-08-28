package logic

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"

	"github.com/louisevanderlith/mango/core/artifact"
)

type InfoHead struct {
	For      string
	ItemID   int64
	ItemName string
}

func GetInfoHead(header string) (InfoHead, error) {
	var result InfoHead
	err := json.Unmarshal([]byte(header), &result)

	return result, err
}

func SaveFile(file multipart.File, header *multipart.FileHeader, info InfoHead) (id int64, err error) {
	var b bytes.Buffer
	copied, err := io.Copy(&b, file)

	if err != nil {
		return -1, err
	}

	blob, mime, err := artifact.NewBLOB(b.Bytes(), info.For)

	if err != nil {
		return -1, err
	}

	upload := artifact.Upload{
		BLOB:     blob,
		Size:     copied,
		Name:     header.Filename,
		ItemID:   info.ItemID,
		ItemName: info.ItemName,
		MimeType: mime,
	}

	rec, err := upload.Create()

	if err != nil {
		return -1, err
	}

	return rec.GetID(), nil
}
