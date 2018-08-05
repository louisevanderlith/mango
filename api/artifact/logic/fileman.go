package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/louisevanderlith/mango/core/artifact"
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
		blob.Data = b.Bytes()

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

			artifact.CreateUpload

			_, err = artifact.Ctx.BLOBs.Create(blob)

			if err == nil {
				id, err = artifact.Ctx.Uploads.Create(&upload)
			}
		}
	}

	return id, err
}
