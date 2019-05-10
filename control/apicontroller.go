package control

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
)

// default paging values
const (
	_page = 1
	_size = 5
)

//APIController serves JSON data.
type APIController struct {
	InstanceController
}

func (ctrl *APIController) Prepare() {
	output := ctrl.Ctx.Output

	output.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	output.Header("Access-Control-Allow-Credentials", "true")
	output.Header("Server", "kettle")
}

func (ctrl *APIController) ServeBinary(data []byte, filename string) {
	dataLen := len(data)
	toTake := 512

	if dataLen < 512 {
		toTake = dataLen
	}

	mimetype := http.DetectContentType(data[:toTake])

	ctrl.ServeBinaryWithMIME(data, filename, mimetype)
}

func (ctrl *APIController) ServeBinaryWithMIME(data []byte, filename, mimetype string) {
	output := ctrl.Ctx.Output

	output.Header("Content-Description", "File Transfer")
	output.Header("Content-Type", mimetype)
	output.Header("Content-Disposition", "attachment; filename="+filename)
	output.Header("Content-Transfer-Encoding", "binary")
	output.Header("Expires", "0")
	output.Header("Cache-Control", "must-revalidate")
	output.Header("Pragma", "public")

	output.Body(data)
}

//Serve sends data as JSON response.
func (ctrl *APIController) Serve(statuscode int, err error, result interface{}) {
	resp := mango.NewRESTResult(statuscode, err, result)

	ctrl.Ctx.Output.SetStatus(resp.Code)

	ctrl.Data["json"] = *resp
	ctrl.ServeJSON()
}

//GetPageData turns /B1 into page 1. size 1
func (ctrl *APIController) GetPageData() (page, pageSize int) {
	pageData := ctrl.Ctx.Input.Param(":pagesize")
	return getPageData(pageData)
}

func getPageData(pageData string) (int, int) {

	if len(pageData) < 2 {
		return _page, _size
	}

	pChar := []rune(pageData[:1])

	if len(pChar) != 1 {
		return _page, _size
	}

	page := int(pChar[0]) % 32
	pageSize, err := strconv.Atoi(pageData[1:])

	if err != nil {
		return _page, _size
	}

	return page, pageSize
}

//GetKeyedRequest will return the Key and update the Target when Requests are sent for updates.
func (ctrl *APIController) GetKeyedRequest(target interface{}) (husk.Key, error) {
	result := struct {
		Key  husk.Key
		Body interface{}
	}{
		Body: target,
	}

	err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &result)

	if err != nil {
		return husk.CrazyKey(), err
	}

	return result.Key, nil
}

//GetMyToken
func (ctrl *APIController) GetMyToken() string {
	vals := ctrl.Ctx.Request.Header["Authorization"]

	if len(vals) == 0 {
		return ""
	}

	return vals[0]
}
