package controllers

import (
	"strconv"

	"github.com/louisevanderlith/mango/api/artifact/logic"
	"github.com/louisevanderlith/mango/core/artifact"
	"github.com/louisevanderlith/mango/util/control"
)

type UploadController struct {
	control.APIController
}

// @Title GetUploads
// @Description Gets the uploads
// @Success 200 {[]artifact.Upload} []artifact.Upload
// @router /:pageData(^[A-Z](?:_?[0-9]+)*$) [get]
func (req *UploadController) Get() {
	page, size := req.GetPageData()

	results, err := artifact.GetUploads(page, size, func(obj artifact.Upload) bool {
		return true
	})

	req.Serve(err, results)
}

// @Title GetUpload
// @Description Gets the requested upload
// @Param	uploadKey			path	husk.Key 	true		"Key of the file you require"
// @Success 200 {artifact.Upload} artifact.Upload
// @router /:uploadID([0-9]+) [get]
func (req *UploadController) GetByID() {
	uploadID, err := husk.() strconv.ParseInt(req.Ctx.Input.Param(":uploadKey"), 10, 64)

	if err != nil {
		req.Serve(err, nil)
	}

	result, err := artifact.GetUpload(uploadKey)

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
		result, filename, err = artifact.GetUploadFile(uploadID)
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
	infoHead, err := logic.GetInfoHead(info)

	if err != nil {
		req.Serve(err, id)
	}

	file, header, err := req.GetFile("file")

	if err != nil {
		req.Serve(err, id)
	}

	defer file.Close()

	id, err = logic.SaveFile(file, header, infoHead)

	req.Serve(err, id)
}
