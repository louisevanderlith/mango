package logic

import (
	"mime/multipart"
	"bytes"
	"bufio"
	"io"
	"github.com/louisevanderlith/mango/db/artifact"
	"github.com/louisevanderlith/mango/util/enums"
	"fmt"
	"errors"
	"github.com/louisevanderlith/mango/db"
)

func SaveFile(file multipart.File, header *multipart.FileHeader) (err error) {
	var b bytes.Buffer
	dst := bufio.NewWriter(&b)
	defer dst.Flush()

	_, err = io.Copy(dst, file)

	if err == nil {
		blob := artifact.Blob{
			Data: b.Bytes(),
		}

		targetType := enums.GetOptimizeType("logo")
		err = blob.OptimizeFor(targetType)
		fmt.Println(header.Header)

		if err == nil {
			upload := artifact.Upload{
				BLOB:     &blob,
				Size:     b.Len(),
				Name:     header.Filename,
				ItemID:   0,
				MimeType: "image/png",
			}

			artifact.Ctx.Upload.Create(&upload)
		}
	}

	return err
}

func GetFile(id int64) (result artifact.Upload, err error) {
	if id > 0 {
		filter := artifact.Upload{}
		filter.ID = id

		var record db.IRecord
		record, err = artifact.Ctx.Upload.ReadOne(&filter)

		result = record.(artifact.Upload)
	} else {
		err = errors.New("ID is invalid.")
	}

	return result, err
}

func getBLOB(id int64) (result artifact.Upload, err error) {
	if id > 0 {
		filter := artifact.Upload{}
		filter.ID = id

		var record db.IRecord
		record, err = artifact.Ctx.Upload.ReadOne(&filter, "BLOB")

		result = record.(artifact.Upload)
	} else {
		err = errors.New("ID is invalid.")
	}

	return result, err
}

func GetFileOnly(id int64) (result []byte, filename string, err error) {
	upload, err := getBLOB(id)

	if err == nil {
		result = upload.BLOB.Data
		filename = upload.Name
	}

	return result, filename, err
}
