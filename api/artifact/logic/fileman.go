package logic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/db/artifact"
	"github.com/louisevanderlith/mango/util/enums"
)

type InfoHead struct {
	For      string
	ItemID   int64
	ItemName string
}

func GetInfoHead(header string) InfoHead {
	var result InfoHead
	err := json.Unmarshal([]byte(header), &result)

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func SaveFile(file multipart.File, header *multipart.FileHeader, info InfoHead) (id int64, err error) {
	var b bytes.Buffer
	_, err = io.Copy(&b, file)

	if err == nil {
		blob := new(artifact.Blob)
		blob.SetData(b.Bytes())

		targetType := enums.GetOptimizeType(info.For)
		err = blob.OptimizeFor(targetType)

		if err == nil {
			upload := artifact.Upload{
				BLOB:     blob,
				Size:     len(blob.Data),
				Name:     header.Filename,
				ItemID:   info.ItemID,
				ItemName: info.ItemName,
				MimeType: "image/png",
			}

			_, err = artifact.Ctx.BLOB.Create(blob)

			if err == nil {
				id, err = artifact.Ctx.Upload.Create(&upload)
			}
		}
	}

	return id, err
}

func GetFile(id int64) (result *artifact.Upload, err error) {
	if id > 0 {
		filter := artifact.Upload{}
		filter.Id = id

		var record db.IRecord
		record, err = artifact.Ctx.Upload.ReadOne(&filter)

		result = record.(*artifact.Upload)
	} else {
		err = errors.New("ID is invalid.")
	}

	return result, err
}

func getUpload(id int64) (result *artifact.Upload, err error) {
	if id > 0 {
		filter := artifact.Upload{}
		filter.Id = id

		var record db.IRecord
		record, err = artifact.Ctx.Upload.ReadOne(&filter, "BLOB")

		result = record.(*artifact.Upload)
	} else {
		err = errors.New("ID is invalid.")
	}

	return result, err
}

func GetFileOnly(id int64) (result []byte, filename string, err error) {
	upload, err := getUpload(id)

	if err == nil {
		result = upload.BLOB.GetData()
		filename = upload.Name
	}

	return result, filename, err
}
