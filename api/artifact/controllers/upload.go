package controllers

import (
	"strconv"

	"github.com/louisevanderlith/mango/api/artifact/logic"
	"github.com/louisevanderlith/mango/db/artifact"
	"github.com/louisevanderlith/mango/util/control"
)

type UploadController struct {
	control.APIController
}

// @Title GetUploads
// @Description Gets the uploads
// @Success 200 {[]artifact.Upload} []artifact.Upload
// @router / [get]
func (req *UploadController) Get() {

	var results artifact.Uploads
	upl := artifact.Upload{}
	err := artifact.Ctx.Uploads.Read(&upl, &results)

	req.Serve(err, results)
}

// @Title GetUpload
// @Description Gets the requested upload
// @Param	uploadID			path	int64 	true		"ID of the file you require"
// @Success 200 {artifact.Upload} artifact.Upload
// @router /:uploadID [get]
func (req *UploadController) GetByID() {
	var result *artifact.Upload
	uploadID, err := strconv.ParseInt(req.Ctx.Input.Param(":uploadID"), 10, 64)

	if err == nil {
		result, err = logic.GetFile(uploadID)
	}

	req.Serve(err, result)
}

// @Title GetFile
// @Description Gets the requested file only
// @Param	uploadID			path	int64 	true		"ID of the file you require"
// @Success 200 {[]byte} []byte
// @router /file/:uploadID [get]
func (req *UploadController) GetFileBytes() {
	var result []byte
	var filename string
	uploadID, err := strconv.ParseInt(req.Ctx.Input.Param(":uploadID"), 10, 64)

	if err == nil {
		result, filename, err = logic.GetFileOnly(uploadID)
	}

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		result = []byte(err.Error())
	}

	req.ServeBinary(result, filename)
}

// @Title UploadFile
// @Description Upload a file
// @Param    file        form     file    true        "File"
// @Param	body		body 	artifact.Upload	true		"body for upload content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *UploadController) Post() {
	var id int64

	info := req.GetString("info")
	infoHead := logic.GetInfoHead(info)
	file, header, err := req.GetFile("file")

	if err == nil {
		defer file.Close()
		id, err = logic.SaveFile(file, header, infoHead)
	}

	req.Serve(err, id)
}
