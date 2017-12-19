package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

type UploadController struct {
	util.SecureController
}

func init() {
	auths := util.ActionAuth{}
	auths["POST"] = enums.Admin

	util.ProtectMethods(auths)
}

func (req *UploadController) Get() {

}

// @Title Upload File
// @Description Upload a file
// @Param    file        form     file    true        "File"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]

func (req *UploadController) Post() {
	req.GetFile()
}
