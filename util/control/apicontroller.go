package control

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/louisevanderlith/mango/util"
)

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

	output.Header("Content-Description", "File Transfer")
	output.Header("Content-Type", "application/octet-stream")
	output.Header("Content-Disposition", "attachment; filename="+filename)
	output.Header("Content-Transfer-Encoding", "binary")
	output.Header("Expires", "0")
	output.Header("Cache-Control", "must-revalidate")
	output.Header("Pragma", "public")

	output.Body(data)
}

func (ctrl *APIController) Serve(result interface{}, err error) {
	resp := util.NewRESTResult(err, result)

	if resp.Failed() {
		ctrl.Ctx.Output.SetStatus(500)
	}

	log.Printf("JSON-%+v\n", resp)
	ctrl.Data["json"] = *resp
	ctrl.ServeJSON()
}

func (ctrl *APIController) GetPageData() (page, pageSize int) {
	pageData := ctrl.Ctx.Input.Param(":pageData")
	page = 0
	pageSize = 10

	if len(pageData) >= 2 {
		pChar, _ := strconv.Atoi(fmt.Sprintf("%c", pageData[0]))

		page = pChar % 32
		pageSize, _ = strconv.Atoi(pageData[1:])
	}

	return page, pageSize
}

func (ctrl *APIController) GetKeyedRequest() (WithKey, error) {
	result := WithKey{}
	err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &result)

	return result, err
}
