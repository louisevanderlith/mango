package controllers

import (
	"github.com/louisevanderlith/mango/api/artifact/logic"
	"strconv"
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

	var results []*artifact.Upload
	upl := artifact.Upload{}
	err := artifact.Ctx.Upload.Read(upl, &results)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": results}
	}

	req.ServeJSON()
}

// @Title GetUpload
// @Description Gets the requested upload
// @Param	uploadID			path	int64 	true		"ID of the file you require"
// @Success 200 {artifact.Upload} artifact.Upload
// @router /:uploadID [get]
func (req *UploadController) GetByID() {
	uploadID, err := strconv.ParseInt(req.Ctx.Input.Param(":uploadID"), 10, 64)

	if err == nil {
		file, err := logic.GetFile(uploadID)

		if err == nil {
			req.Data["json"] = map[string]interface{}{"Data": file}
		}
	}

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	}

	req.ServeJSON()
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
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *UploadController) Post() {
	file, header, err := req.GetFile("file")

	if err != nil {
		defer file.Close()
		err = logic.SaveFile(file, header)
	}

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "File Saved"}
	}

	req.ServeJSON()
}
