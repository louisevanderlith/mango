package control

import (
	"encoding/json"
	"log"
	"strconv"

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
	output.Header("Server", "kettle")
}

func (ctrl *APIController) ServeBinary(data []byte, filename string) {
	output := ctrl.Ctx.Output

	mimetype := http.DetectContentType(data[:512])

	output.Header("Content-Description", "File Transfer")
	output.Header("Content-Type", mimetype)
	output.Header("Content-Disposition", "attachment; filename="+filename)
	output.Header("Content-Transfer-Encoding", "binary")
	output.Header("Expires", "0")
	output.Header("Cache-Control", "must-revalidate")
	output.Header("Pragma", "public")

	output.Body(data)
}

func (ctrl *APIController) Serve(result interface{}, err error) {
	resp := mango.NewRESTResult(err, result)

	if resp.Failed() {
		log.Printf("\t [API Error]: %s\n", resp.Reason)
		ctrl.Ctx.Output.SetStatus(500)
	}

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
	//pChar, err := strconv.Atoi(pageData[:1][0])

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

func (ctrl *APIController) GetKeyedRequest() (WithKey, error) {
	result := WithKey{}
	err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &result)

	return result, err
}

func GetFileContentType(firstBits []byte) (string, error) {

    // Only the first 512 bytes are used to sniff the content type.
    buffer := make([]byte, 512)

    _, err := out.Read(buffer)
    if err != nil {
        return "", err
    }

    // Use the net/http package's handy DectectContentType function. Always returns a valid 
    // content-type by returning "application/octet-stream" if no others seemed to match.
    contentType := 

    return contentType, nil
}